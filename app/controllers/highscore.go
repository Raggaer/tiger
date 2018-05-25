package controllers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

var highscoreCommand = Command{
	Usage: "top option",
	Options: []CommandOption{
		{
			Name:        "experience",
			Description: "Returns server highest level players",
		},
		{
			Name:        "magic",
			Description: "Returns server highest magic level players",
		},
		{
			Name:        "fist",
			Description: "Returns server highest fist fighting players",
		},
	},
}

// ViewTopPlayersSkillFist returns a list of the server top skill fist players
func ViewTopPlayersSkillFist(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillFist(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_fist", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by fist fighting",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersMagicLevel returns a list of the server top magic level players
func ViewTopPlayersMagicLevel(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersByMagicLevel(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_maglevel", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by magic level",
		Description: data,
		Color:       3447003,
	}, nil
}

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
