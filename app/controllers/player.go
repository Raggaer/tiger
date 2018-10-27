package controllers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
	"github.com/schollz/closestmatch"
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

func playerNotFound(context *Context, player string) (*discordgo.MessageEmbed, error) {
	playerFuzzy, ok := context.Cache.Get("player_fuzzy")
	if ok {
		c, ok := playerFuzzy.(*closestmatch.ClosestMatch)
		if !ok {
			return nil, errors.New("Wrong players fuzzy cache")
		}

		// Retrieve possible related word
		w1 := c.Closest(player)
		msg := ""
		if w1 == "" {
			msg = "Player not found"
		} else {
			msg = fmt.Sprintf("Player not found. Maybe you wanted to say **%s**", w1)
		}

		// Render player not found message
		return &discordgo.MessageEmbed{
			Title:       "Player  not found",
			Description: msg,
			Color:       3447003,
		}, nil
	}

	// Gather all players and create closest match
	players, err := models.GetPlayersFuzzy(context.DB)
	if err != nil {
		return nil, err
	}

	c := closestmatch.New(players, []int{2})
	context.Cache.Set("player_fuzzy", c, time.Minute*5)
	w1 := c.Closest(player)
	msg := ""
	if w1 == "" {
		msg = "Player not found"
	} else {
		msg = fmt.Sprintf("Player not found. Maybe you wanted to say **%s**", w1)
	}

	// Render player not found message
	return &discordgo.MessageEmbed{
		Title:       "Player  not found",
		Description: msg,
		Color:       3447003,
	}, nil
}

// ViewPlayer views the given server character
func ViewPlayer(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(m.Content))
	if err != nil {
		return playerNotFound(context, strings.TrimSpace(m.Content))
	}

	// Retrieve player vocation
	playerVocation := "No vocation"
	for _, v := range context.Vocations {
		if v.ID == player.Vocation {
			playerVocation = v.Name
			break
		}
	}

	return context.ExecuteTemplate("player_info", map[string]interface{}{
		"vocationName": playerVocation,
		"player":       player,
	})
}

// ViewPlayerDeaths retrieves the last deaths of the given player
func ViewPlayerDeaths(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(m.Content))
	if err != nil {
		return playerNotFound(context, strings.TrimSpace(m.Content))
	}

	// Retrieve player deaths
	deaths, err := models.GetPlayerDeaths(context.DB, player, 10)
	if err != nil {
		return nil, err
	}

	return context.ExecuteTemplate("player_death", map[string]interface{}{
		"deaths": deaths,
		"player": player,
	})
}
