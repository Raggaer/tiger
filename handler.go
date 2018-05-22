package main

import (
	"database/sql"
	"strings"
	"time"

	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/controllers"
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
	handlers.Add("vocation", controllers.ViewVocation)
	handlers.Add("version", controllers.Version)
	handlers.Add("uptime", controllers.Uptime)
	handlers.Add("monster", controllers.ViewMonster)

}

// Add registers a new handler
func (h *handlerList) Add(prefix string, hd interface{}) {
	h.list = append(h.list, handler{
		Prefix:  h.cfg.Discord.Prefix + prefix,
		Handler: hd,
	})
}

func handleCreateMessage(cfg *config.Config, tasks *xmlTaskList, db *sql.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Create controller context
	ctx := controllers.Context{
		Config:    cfg,
		Monsters:  tasks.Monsters,
		Items:     tasks.Items,
		Vocations: tasks.Vocations,
		Start:     time.Now(),
		DB:        db,
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

		// Loop all registered handlers
		for _, h := range handlers.list {
			if strings.HasPrefix(m.Content, h.Prefix) {

				// Check if we can execute the handler
				handlerFunc, ok := h.Handler.(func(*controllers.Context, *discordgo.Session, *discordgo.MessageCreate) error)
				if !ok {
					continue
				}

				// Remove prefix from content
				m.Content = strings.TrimSpace(strings.TrimPrefix(m.Content, h.Prefix))

				// Execute handler
				if err := handlerFunc(&ctx, s, m); err != nil {
					log.Printf("Unable to execute handlerfunc %s: %v", h.Prefix, err)
				}
				break
			}
		}
	}
}
