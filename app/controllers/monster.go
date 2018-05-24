package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
	"github.com/raggaer/tiger/app/xml"
)

var monsterCommand = Command{
	Usage: "monster Name, option",
	Options: []CommandOption{
		{
			Name:        "info",
			Description: "Returns basic information about the monster",
		},
		{
			Name:        "loot",
			Description: "Returns the loot table of the monster",
		},
		{
			Name:        "deaths",
			Description: "Returns the last 10 players killed by the monster",
		},
	},
}

// MonsterLoot defines a list of monster loot
type MonsterLoot struct {
	Item   string
	Chance float64
}

// ViewMonster returns information about the given monster
func ViewMonster(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := strings.Split(m.Content, ",")
	if len(data) <= 1 {
		return monsterCommand.RenderUsage("Unknown option", context, s, m)
	}

	// Get monster
	monster, ok := context.Monsters[strings.ToLower(strings.TrimSpace(data[0]))]
	if !ok {
		return monsterCommand.RenderUsage("Monster not found", context, s, m)
	}

	// Switch monster view method
	switch strings.TrimSpace(data[1]) {
	case "loot":
		return viewMonsterLoot(context, s, m, monster)
	case "info":
		return viewMonsterInformation(context, s, m, monster)
	case "deaths":
		return viewMonsterKilledPlayers(context, s, m, monster)
	default:
		return viewMonsterInformation(context, s, m, monster)
	}
}

func viewMonsterKilledPlayers(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	// Load monster deaths
	deaths, err := models.GetPlayerDeathsByMonster(context.DB, monster, 10)
	if err != nil {
		return err
	}

	data, err := context.ExecuteTemplate("monster_death.md", map[string]interface{}{
		"deaths":  deaths,
		"monster": monster,
	})
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Players killed by " + monster.Description,
		Description: data,
	})
	return err
}

func viewMonsterInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	data, err := context.ExecuteTemplate("monster_info.md", map[string]interface{}{
		"monster": monster,
	})
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Information about " + monster.Description,
		Description: data,
	})
	return err
}

func viewMonsterLoot(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	// Create loot list
	loot := []*MonsterLoot{}
	for _, item := range monster.Loot.Loot {

		// Calculate item chance percentage
		chance := 100.0
		if item.Chance > 0 {
			chance = float64(item.Chance) / 1000.0
		}

		// Check if we need to retrieve item from map
		if item.Name == "" {
			i, ok := context.Items[item.ID]
			if !ok {
				continue
			}
			loot = append(loot, &MonsterLoot{
				Item:   i.Name,
				Chance: chance,
			})
			continue
		}

		loot = append(loot, &MonsterLoot{
			Item:   item.Name,
			Chance: chance,
		})
	}

	data, err := context.ExecuteTemplate("monster_loot.md", map[string]interface{}{
		"monster": monster,
		"loot":    loot,
	})
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       monster.Name + " loot",
		Description: data,
	})
	return err
}
