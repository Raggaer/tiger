package controllers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

// ViewTopPlayersExperience returns a list of the server top experience players
func ViewTopPlayersExperience(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersByExperience(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_experience", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by experience",
		Description: data,
		Color:       3447003,
	}, nil
}
