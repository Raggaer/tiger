package controllers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var vocationCommand = Command{
	Usage:       "vocation name",
	Description: "Provides information about server vocations",
}

// ViewVocation sends information about the given vocation
func ViewVocation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve vocation by name
	voc, ok := context.Vocations[strings.ToLower(m.Content)]
	if !ok {
		return vocationCommand.RenderUsage("Vocation not found", context, s, m)
	}

	return context.ExecuteTemplate("vocation_info", map[string]interface{}{
		"voc": voc,
	})
}
