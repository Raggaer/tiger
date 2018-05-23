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
	_, err := models.GetPlayerByName(context.DB, strings.TrimSpace(data[0]))
	if err != nil {
		return viewPlayerCommand.RenderUsage("Player not found", context, s, m)
	}

	switch strings.TrimSpace(data[1]) {
	case "info":
	}

	// Create player message
	return nil
}

func viewPlayerInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, player *models.Player) {

}
