package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/models"
)

func monitorServerPlayerDeaths(guild *discordgo.Guild, cfg *config.Config, tick time.Duration, db *sql.DB, s *discordgo.Session) {
	// Create event ticker
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	// Retrieve valid death channels
	deathChannels := retrieveValidChannels(guild, cfg.Discord.DeathChannels)

	// Wait for ticker channel
	for t := range ticker.C {
		deaths, err := models.GetTimeServerDeaths(db, 10, t)
		if err != nil {
			continue
		}

		// Create discord message
		for _, ch := range deathChannels {
			for _, death := range deaths {
				s.ChannelMessageSendEmbed(ch, &discordgo.MessageEmbed{
					Title: "Player death",
					Color: 3447003,
					Description: fmt.Sprintf(
						"Player **%s** killed by **%s**",
						death.Player.Name,
						death.KilledBy,
					),
				})
			}
		}
	}
}

func retrieveValidChannels(guild *discordgo.Guild, channelNames []string) []string {
	validChannels := []string{}
	for _, ch := range guild.Channels {
		for _, dh := range channelNames {
			if dh == ch.Name {
				validChannels = append(validChannels, ch.ID)
			}
		}
	}
	return validChannels
}
