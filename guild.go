package main

import (
	"database/sql"
	"text/template"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	cache "github.com/robfig/go-cache"
)

func handleGuildCreate(cfg *config.Config, tasks *xmlTaskList, db *sql.DB, tpl *template.Template, cache *cache.Cache) func(s *discordgo.Session, event *discordgo.GuildCreate) {
	return func(s *discordgo.Session, event *discordgo.GuildCreate) {
		go monitorServerPlayerDeaths(event.Guild, cfg, time.Minute, db, s)
	}
}
