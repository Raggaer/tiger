package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

var monsterCommand = Command{
	Usage: "monster option name",
	Options: []CommandOption{
		{
			Name:        "view",
			Description: "Returns basic information about the monster",
		},
		{
			Name:        "loot",
			Description: "Returns the loot table of the monster",
		},
		{
			Name:        "victims",
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
func ViewMonster(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Get monster
	monster, ok := context.Monsters[strings.ToLower(m.Content)]
	if !ok {
		return monsterCommand.RenderUsage("Monster not found", context, s, m)
	}

	data, err := context.ExecuteTemplate("monster_info", map[string]interface{}{
		"monster": monster,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Information about " + monster.Description,
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewMonsterKilledPlayers returns the list of victims od the given monster
func ViewMonsterKilledPlayers(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Get monster
	monster, ok := context.Monsters[strings.ToLower(m.Content)]
	if !ok {
		return monsterCommand.RenderUsage("Monster not found", context, s, m)
	}

	// Load monster deaths
	deaths, err := models.GetPlayerDeathsByMonster(context.DB, monster, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("monster_death", map[string]interface{}{
		"deaths":  deaths,
		"monster": monster,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Players killed by " + monster.Description,
		Description: data,
		Color:       3447003,
	}, nil
}

func ViewMonsterLoot(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Get monster
	monster, ok := context.Monsters[strings.ToLower(m.Content)]
	if !ok {
		return monsterCommand.RenderUsage("Monster not found", context, s, m)
	}

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

	data, err := context.ExecuteTemplate("monster_loot", map[string]interface{}{
		"monster": monster,
		"loot":    loot,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       monster.Name + " loot",
		Description: data,
		Color:       3447003,
	}, nil
}
