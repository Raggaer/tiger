package main

import (
	"database/sql"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/controllers"
	"github.com/raggaer/tiger/app/xml"
	cache "github.com/robfig/go-cache"
)

func handleGuildCreate(cfg *config.Config, tasks *xmlTaskList, db *sql.DB, tpl map[string]*xml.CommandTemplate, cache *cache.Cache) func(s *discordgo.Session, event *discordgo.GuildCreate) {
	return func(s *discordgo.Session, event *discordgo.GuildCreate) {
		// Create context
		ctx := &controllers.Context{
			Template: tpl,
			Config:   cfg,
			Cache:    cache,
			DB:       db,
		}
		go monitorServerPlayerDeaths(event.Guild, time.Second*5, ctx, s, event)
	}
}
