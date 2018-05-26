package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
	lua "github.com/yuin/gopher-lua"
)

// Config defines the application config file
type Config struct {
	Template templateConfig
	Server   serverConfig
	Discord  discordConfig
	Database databaseConfig
}

type templateConfig struct {
	Directory string
	Extension string
}

type serverConfig struct {
	Path string
}

type discordConfig struct {
	Token    string
	Prefix   string
	Status   string
	Channels []string
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
	Schema   string
}

// Load loads the application config file
func Load(path string) (*Config, error) {
	cfg := Config{}

	// Decode config.toml config file
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	// Check if database is already set
	if cfg.Database.Schema == "" {
		if err := loadServerConfig(filepath.Join(cfg.Server.Path, "config.lua"), &cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

func loadServerConfig(path string, cfg *Config) error {
	state := lua.NewState()
	defer state.Close()

	// Load server config.lua file
	if err := state.DoFile(path); err != nil {
		return err
	}

	// Set database values into config struct
	cfg.Database.Host = string(state.GetGlobal("mysqlHost").(lua.LString))
	cfg.Database.User = string(state.GetGlobal("mysqlUser").(lua.LString))
	cfg.Database.Password = string(state.GetGlobal("mysqlPass").(lua.LString))
	cfg.Database.Schema = string(state.GetGlobal("mysqlDatabase").(lua.LString))

	return nil
}
