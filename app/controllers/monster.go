package controllers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/xml"
)

// ViewMonster returns information about the given monster
func ViewMonster(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := strings.Split(m.Content, ",")
	if len(data) <= 1 {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Color:       3447003,
			Title:       "View monster",
			Description: "Not enough parameters. Command usage `" + context.Config.Discord.Prefix + "monster Name, option`",
		})
		return err
	}

	// Get monster
	monster, ok := context.Monsters[strings.TrimSpace(data[0])]
	if !ok {
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Color:       3447003,
			Title:       "Monster not found",
			Description: "Monster not found. Command usage `" + context.Config.Discord.Prefix + "monster Name, option`",
		})
		return err
	}

	// Switch monster view method
	switch strings.TrimSpace(data[1]) {
	case "loot":
		return viewMonsterLoot(context, s, m, monster)
	case "info":
		return viewMonsterInformation(context, s, m, monster)
	}
	return nil
}

func viewMonsterInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	msg := "**You view " + monster.Description + "** \r\n \r\n"
	msg += fmt.Sprintf("- **Experience**: %d \r\n", monster.Experience)
	msg += fmt.Sprintf("- **Speed**: %d \r\n", monster.Speed)
	msg += fmt.Sprintf("- **Health**: %d \r\n \r\n", monster.Health.Now)
	for _, attack := range monster.Attacks.Attacks {
		msg += fmt.Sprintf("- **Attack**: %s (%d, %d)\r\n", attack.Name, attack.Min, attack.Max)
	}
	// Send information message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       monster.Name + " information",
		Description: msg,
	})
	return err
}

func viewMonsterLoot(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	msg := ""
	for _, item := range monster.Loot.Loot {
		// Calculate item chance percentage
		chance := 100.0
		if item.Chance > 0 {
			chance = float64(item.Chance) / 1000.0
		}
		// Check if we need to retrieve item from map
		if item.Name == "" {
			i, ok := context.Items[item.ID]
			if !ok {
				continue
			}
			msg += fmt.Sprintf("\r\n- **%s** - %.2f", i.Name, chance) + "%"
			continue
		}
		msg += fmt.Sprintf("\r\n- **%s** - %.2f", item.Name, chance) + "%"
	}

	// Send loot message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       monster.Name + " loot",
		Description: msg,
	})
	return err
}