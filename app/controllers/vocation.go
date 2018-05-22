package controllers

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ViewVocation sends information about the given vocation
func ViewVocation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Retrieve vocation by name
	voc, ok := context.Vocations[m.Content]
	if !ok {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Color:       3447003,
			Title:       "Vocation not found",
			Description: "Vocation not found. Commando usage `" + context.Config.Discord.Prefix + "vocation Name`",
		})
		return err
	}

	// Write message with vocation information
	msg := &strings.Builder{}
	msg.WriteString("**You see a " + voc.Description + "** \r\n \r\n")
	msg.WriteString("- **Gain capacity**: " + strconv.Itoa(voc.GainCap))
	msg.WriteString("\r\n- **Gain health**: " + strconv.Itoa(voc.GainHealth))
	msg.WriteString("\r\n- **Gain mana**: " + strconv.Itoa(voc.GainMana))
	msg.WriteString("\r\n- **Base speed**: " + strconv.Itoa(voc.BaseSpeed))
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Vocation " + voc.Name,
		Description: msg.String(),
	})
	return err
}