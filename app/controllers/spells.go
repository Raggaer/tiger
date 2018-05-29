package controllers

import "github.com/bwmarrin/discordgo"

// ViewSpell returns information about the given spell
func ViewSpell(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	var data string
	var err error

	instantSpell, ok := context.InstantSpells[m.Content]
	if ok {
		data, err = context.ExecuteTemplate("view_instant", map[string]interface{}{
			"spell": instantSpell,
		})
		if err != nil {
			return nil, err
		}
	}

	conjureSpell, ok := context.ConjureSpells[m.Content]
	if ok {
		data, err = context.ExecuteTemplate("view_conjure", map[string]interface{}{
			"conjure": conjureSpell,
		})
		if err != nil {
			return nil, err
		}
	}

	return &discordgo.MessageEmbed{
		Title:       "View spell " + m.Content,
		Description: data,
		Color:       3447003,
	}, nil
}
