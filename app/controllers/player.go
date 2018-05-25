package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

var viewPlayerCommand = Command{
	Usage: "player option name",
	Options: []CommandOption{
		{
			Name:        "view",
			Description: "Shows the basic information about a character",
		},
		{
			Name:        "deaths",
			Description: "Latest character deaths",
		},
	},
}

// ViewPlayer views the given server character
func ViewPlayer(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(m.Content))
	if err != nil {
		return viewPlayerCommand.RenderUsage("Player not found", context, s, m)
	}

	// Retrieve player vocation
	playerVocation := "No vocation"
	for _, v := range context.Vocations {
		if v.ID == player.Vocation {
			playerVocation = v.Name
			break
		}
	}

	data, err := context.ExecuteTemplate("player_info", map[string]interface{}{
		"vocationName": playerVocation,
		"player":       player,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "View player " + player.Name,
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewPlayerDeaths retrieves the last deaths of the given player
func ViewPlayerDeaths(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(m.Content))
	if err != nil {
		return viewPlayerCommand.RenderUsage("Player not found", context, s, m)
	}

	// Retrieve player deaths
	deaths, err := models.GetPlayerDeaths(context.DB, player, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("player_death", map[string]interface{}{
		"deaths": deaths,
		"player": player,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Latest " + player.Name + " deaths",
		Description: data,
		Color:       3447003,
	}, nil
}
