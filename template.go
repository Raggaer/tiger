package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/controllers"
	"github.com/raggaer/tiger/app/xml"
)

func reloadTemplates(context *controllers.Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Check if user is administrator
	perms, err := s.State.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return nil, err
	}

	// Check if user has administrator permission
	if perms&discordgo.PermissionAdministrator <= 0 {
		return nil, nil
	}

	// Reload templates
	tpl, err := loadTemplates(context.Config.Template.Directory, context.Config.Template.Extension)
	if err != nil {
		return nil, err
	}
	context.Template = tpl

	return &discordgo.MessageEmbed{
		Title:       "Reload templates",
		Description: "Templates loaded",
		Color:       3447003,
	}, nil
}

func loadTemplates(path, extension string) (map[string]*xml.CommandTemplate, error) {
	list := map[string]*xml.CommandTemplate{}
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), extension) {
			tpl, err := xml.ParseTemplate(path)
			if err != nil {
				return err
			}
			list[info.Name()] = tpl
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return list, nil
}
