package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/models"
)

func monitorServerPlayerDeaths(cfg *config.Config, tick time.Duration, db *sql.DB, s *discordgo.Session) {
	// Create event ticker
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	// Get valid death channels from config
	guid := s.State.Guilds[0]
	guild, err := s.Guild(guid.ID)
	if err != nil {
		log.Fatalf("Unable to retrieve guild channels: %v", err)
	}

	// Wait for the guild channels to be ready
	for len(guild.Channels) <= 0 {
		time.Sleep(time.Second)
	}

	deathChannels := []string{}
	for _, ch := range guild.Channels {
		for _, dh := range cfg.Discord.DeathChannels {
			if dh == ch.Name {
				deathChannels = append(deathChannels, ch.ID)
			}
		}
	}

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
