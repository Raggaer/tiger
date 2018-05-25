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
		{
			Name:        "club",
			Description: "Returns server highest club fighting players",
		},
		{
			Name:        "axe",
			Description: "Returns server highest axe fighting players",
		},
		{
			Name:        "sword",
			Description: "Returns server highest sword fighting players",
		},
		{
			Name:        "dist",
			Description: "Returns server highest distance fighting players",
		},
		{
			Name:        "shield",
			Description: "Returns server highest shielding players",
		},
		{
			Name:        "fish",
			Description: "Returns server highest fishing players",
		},
	},
}

// ViewTopPlayersSkillFishing returns a list of the server top skill fishing players
func ViewTopPlayersSkillFishing(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillFishing(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_fishing", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by fishing",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersSkillShielding returns a list of the server top skill shielding players
func ViewTopPlayersSkillShielding(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillShield(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_shield", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by shielding",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersSkillDist returns a list of the server top skill dist players
func ViewTopPlayersSkillDist(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillDist(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_dist", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by distance fighting",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersSkillAxe returns a list of the server top skill axe players
func ViewTopPlayersSkillAxe(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillAxe(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_axe", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by axe fighting",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersSkillSword returns a list of the server top skill sword players
func ViewTopPlayersSkillSword(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillSword(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_sword", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by sword fighting",
		Description: data,
		Color:       3447003,
	}, nil
}

// ViewTopPlayersSkillClub returns a list of the server top skill club players
func ViewTopPlayersSkillClub(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve top players
	players, err := models.GetTopPlayersBySkillClub(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("top_club", map[string]interface{}{
		"players": players,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Top players by club fighting",
		Description: data,
		Color:       3447003,
	}, nil
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
