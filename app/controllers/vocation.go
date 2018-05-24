package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var vocationCommand = Command{
	Usage:       "vocation Name",
	Description: "Provides information about server vocations",
}

// ViewVocation sends information about the given vocation
func ViewVocation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Retrieve vocation by name
	voc, ok := context.Vocations[strings.ToLower(m.Content)]
	if !ok {
		return vocationCommand.RenderUsage("Vocation not found", context, s, m)
	}

	data, err := context.ExecuteTemplate("vocation_info.md", map[string]interface{}{
		"voc": voc,
	})
	if err != nil {
		return err
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Vocation " + voc.Name,
		Description: data,
	})
	return err
}
