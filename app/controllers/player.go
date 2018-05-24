package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

var viewPlayerCommand = Command{
	Usage: "player Name, option",
	Options: []CommandOption{
		{
			Name:        "info",
			Description: "Shows the basic information about a character",
		},
	},
}

// ViewPlayer views the given server character
func ViewPlayer(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := strings.Split(m.Content, ",")
	if len(data) <= 1 {
		return viewPlayerCommand.RenderUsage("Unknown option", context, s, m)
	}

	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(data[0]))
	if err != nil {
		return viewPlayerCommand.RenderUsage("Player not found", context, s, m)
	}

	switch strings.TrimSpace(data[1]) {
	case "info":
		return viewPlayerInformation(context, s, m, player)
	case "deaths":
		return viewPlayerDeaths(context, s, m, player)
	default:
		return viewPlayerInformation(context, s, m, player)
	}
}

func viewPlayerDeaths(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, player *models.Player) error {
	// Retrieve player deaths
	deaths, err := models.GetPlayerDeaths(context.DB, player, 10)
	if err != nil {
		return err
	}

	data, err := context.ExecuteTemplate("player_death.md", map[string]interface{}{
		"deaths": deaths,
		"player": player,
	})
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Latest " + player.Name + " deaths",
		Color:       3447003,
		Description: data,
	})
	return nil
}

func viewPlayerInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, player *models.Player) error {
	// Retrieve player vocation
	playerVocation := "No vocation"
	for _, v := range context.Vocations {
		if v.ID == player.Vocation {
			playerVocation = v.Name
			break
		}
	}

	data, err := context.ExecuteTemplate("player_info.md", map[string]interface{}{
		"vocationName": playerVocation,
		"player":       player,
	})
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "View player " + player.Name,
		Color:       3447003,
		Description: data,
	})
	return err
}
