package controllers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// ViewSpell returns information about the given spell
func ViewSpell(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	instantSpell, ok := context.InstantSpells[m.Content]
	if ok {
		return context.ExecuteTemplate("view_instant", map[string]interface{}{
			"spell": instantSpell,
		})
	}

	conjureSpell, ok := context.ConjureSpells[m.Content]
	if ok {
		return context.ExecuteTemplate("view_conjure", map[string]interface{}{
			"conjure": conjureSpell,
		})
	}

	// Retrieve possible related word
	w1 := context.InstantSpellsFuzzySearch.Closest(m.Content)
	w2 := context.ConjureSpellsFuzzySearch.Closest(m.Content)
	msg := ""
	if w1 != "" && w2 != "" {
		msg = fmt.Sprintf("Maybe you wanted to say **%s** or **%s**", w1, w2)
	} else if w1 == "" && w2 != "" {
		msg = fmt.Sprintf("Maybe you wanted to say **%s**", w2)
	} else {
		msg = fmt.Sprintf("Maybe you wanted to say **%s**", w1)
	}

	// Render spell not found message
	return &discordgo.MessageEmbed{
		Title:       "Spell not found " + m.Content,
		Description: "Spell not found. " + msg,
		Color:       3447003,
	}, nil
}
