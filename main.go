package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
)

func main() {
	// Retrieve execution path
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Unable to retrieve executable path: %v", err)
	}

	// Load application config
	cfg, err := config.Load(filepath.Join(filepath.Dir(execPath), "config.toml"))
	if err != nil {
		log.Fatalf("Unable to load application config: %v", err)
	}

	// Open database connection
	db, err := loadDatabase(cfg.Database.User, cfg.Database.Password, cfg.Database.Schema)
	if err != nil {
		log.Fatalf("Unable to open database connection: %v", err)
	}

	// Register handlers
	registerHandlers(cfg)

	// Load server data
	tasks, taskErr := loadServerData(cfg)
	if taskErr != nil {
		log.Fatalf("Unable to complete xml task %s: %v", taskErr.Name, taskErr.Error)
	}

	// Create discord session
	dg, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		log.Fatalf("Unable to create discord session: %v", err)
	}

	// Load templates
	tpl, err := loadTemplates(cfg.Template.Directory, cfg.Template.Extension)
	if err != nil {
		log.Fatalf("Unable to load template files: %v", err)
	}

	// Register mesasge handler
	dg.AddHandler(handleCreateMessage(cfg, tasks, db, tpl))

	// Open discord session
	if err := dg.Open(); err != nil {
		log.Fatalf("Unable to open discord session: %v", err)
	}
	if err := dg.UpdateStatus(0, cfg.Discord.Status); err != nil {
		log.Fatalf("Unable to set bot status: %v", err)
	}

	// Wait here until CTRL-C or other term signal is received
	log.Println("Tiger is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
