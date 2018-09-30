package main

import (
	"database/sql"
	"strings"
	"text/template"
	"time"

	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/controllers"
	cache "github.com/robfig/go-cache"
)

type handlerList struct {
	list []handler
	cfg  *config.Config
}

type handler struct {
	Prefix  string
	Handler interface{}
}

var handlers handlerList

func registerHandlers(cfg *config.Config) {
	// Create handler object
	handlers = handlerList{
		cfg: cfg,
	}

	// Register handlers
	handlers.Add("spell", controllers.ViewSpell)
	handlers.Add("top fish", controllers.ViewTopPlayersSkillFishing)
	handlers.Add("top shield", controllers.ViewTopPlayersSkillShielding)
	handlers.Add("top dist", controllers.ViewTopPlayersSkillDist)
	handlers.Add("top axe", controllers.ViewTopPlayersSkillAxe)
	handlers.Add("top sword", controllers.ViewTopPlayersSkillSword)
	handlers.Add("top club", controllers.ViewTopPlayersSkillClub)
	handlers.Add("top fist", controllers.ViewTopPlayersSkillFist)
	handlers.Add("top magic", controllers.ViewTopPlayersMagicLevel)
	handlers.Add("top experience", controllers.ViewTopPlayersExperience)
	handlers.Add("vocation", controllers.ViewVocation)
	handlers.Add("version", controllers.Version)
	handlers.Add("uptime", controllers.Uptime)
	handlers.Add("monster view", controllers.ViewMonster)
	handlers.Add("monster victims", controllers.ViewMonsterKilledPlayers)
	handlers.Add("monster loot", controllers.ViewMonsterLoot)
	handlers.Add("player view", controllers.ViewPlayer)
	handlers.Add("player deaths", controllers.ViewPlayerDeaths)
	handlers.Add("deaths", controllers.LatestDeaths)
	handlers.Add("status", controllers.ServerStatus)
	handlers.Add("about", controllers.About)
	handlers.Add("reload templates", reloadTemplates)
}

// Add registers a new handler
func (h *handlerList) Add(prefix string, hd interface{}) {
	h.list = append(h.list, handler{
		Prefix:  h.cfg.Discord.Prefix + prefix,
		Handler: hd,
	})
}

func handleCreateMessage(cfg *config.Config, tasks *xmlTaskList, db *sql.DB, tpl *template.Template, cache *cache.Cache) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Create controller context
	ctx := controllers.Context{
		Config:                   cfg,
		Monsters:                 tasks.Monsters,
		Items:                    tasks.Items,
		Vocations:                tasks.Vocations,
		InstantSpells:            tasks.InstantSpells,
		InstantSpellsFuzzySearch: tasks.InstantSpellsFuzzySearch,
		RuneSpells:               tasks.RuneSpells,
		RuneSpellsFuzzySearch:    tasks.RuneSpellsFuzzySearch,
		ConjureSpells:            tasks.ConjureSpells,
		ConjureSpellsFuzzySearch: tasks.ConjureSpellsFuzzySearch,
		Start:    time.Now(),
		DB:       db,
		Template: tpl,
		Cache:    cache,
	}

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore all messages created by the bot
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Ignore bot messages
		if m.Author.Bot {
			return
		}

		// Check if valid channel
		channelName, err := getChannelName(s, m.ChannelID)
		if err != nil {
			log.Fatalf("Unable to retrieve channel name: %v", err)
		}
		validChannel := false
		for _, n := range cfg.Discord.Channels {
			if n == channelName {
				validChannel = true
				break
			}
		}
		if !validChannel {
			return
		}

		// Loop all registered handlers
		for _, h := range handlers.list {
			if strings.HasPrefix(m.Content, h.Prefix) {
				// Check if we can execute the handler
				handlerFunc, ok := h.Handler.(func(*controllers.Context, *discordgo.Session, *discordgo.MessageCreate) (*discordgo.MessageEmbed, error))
				if !ok {
					continue
				}

				// Remove prefix from content
				m.Content = strings.TrimSpace(strings.TrimPrefix(m.Content, h.Prefix))

				// Create working message
				workMessage, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Color:       3447003,
					Title:       "Working",
					Description: "Working...",
				})
				if err != nil {
					log.Fatalf("Unable to send working message: %v", err)
				}

				// Execute handler
				data, err := handlerFunc(&ctx, s, m)
				if err != nil {
					log.Printf("Unable to execute handlerfunc %s: %v", h.Prefix, err)
					break
				}

				// Delete message if there is nothing to show
				if data == nil {
					s.ChannelMessageDelete(workMessage.ChannelID, workMessage.ID)
				} else {
					// Edit working message
					s.ChannelMessageEditEmbed(workMessage.ChannelID, workMessage.ID, data)
				}
				break
			}
		}
	}
}
