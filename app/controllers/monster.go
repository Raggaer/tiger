package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
	"github.com/raggaer/tiger/app/xml"
)

var monsterCommand = Command{
	Usage: "monster Name, option",
	Options: []CommandOption{
		{
			Name:        "info",
			Description: "Returns basic information about the monster",
		},
		{
			Name:        "loot",
			Description: "Returns the loot table of the monster",
		},
		{
			Name:        "deaths",
			Description: "Returns the last 10 players killed by the monster",
		},
	},
}

// ViewMonster returns information about the given monster
func ViewMonster(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := strings.Split(m.Content, ",")
	if len(data) <= 1 {
		return monsterCommand.RenderUsage("Unknown option", context, s, m)
	}

	// Get monster
	monster, ok := context.Monsters[strings.ToLower(strings.TrimSpace(data[0]))]
	if !ok {
		return monsterCommand.RenderUsage("Monster not found", context, s, m)
	}

	// Switch monster view method
	switch strings.TrimSpace(data[1]) {
	case "loot":
		return viewMonsterLoot(context, s, m, monster)
	case "info":
		return viewMonsterInformation(context, s, m, monster)
	case "deaths":
		return viewMonsterKilledPlayers(context, s, m, monster)
	default:
		return viewMonsterInformation(context, s, m, monster)
	}
}

func viewMonsterKilledPlayers(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	deaths, err := models.GetPlayerDeathsByMonster(context.DB, monster, 10)
	if err != nil {
		return err
	}

	// Render deaths message
	msg := strings.Builder{}
	if len(deaths) <= 0 {
		msg.WriteString("\r\nNoone")
	} else {
		for i, d := range deaths {
			deathTime := timeAgo(time.Unix(d.Time, 0), time.Now())
			msg.WriteString("\r\n" + strconv.Itoa(i+1) + ". Killed **" + d.Player.Name + "** at level **" + strconv.Itoa(d.Level) + "** - *" + deathTime + " ago*")
		}
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Players killed by " + monster.Description,
		Description: msg.String(),
	})
	return err
}

func viewMonsterInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	msg := strings.Builder{}
	msg.WriteString("**You view" + monster.Description + "** \r\n \r\n")
	msg.WriteString("- **Experience**: " + strconv.Itoa(monster.Experience) + "\r\n")
	msg.WriteString("- **Speed**: " + strconv.Itoa(monster.Speed) + " \r\n")
	msg.WriteString("- **Health**: " + strconv.Itoa(monster.Health.Now) + " \r\n \r\n")
	for _, attack := range monster.Attacks.Attacks {
		msg.WriteString("- **Attack**: " + attack.Name + " (" + strconv.Itoa(attack.Min) + ", " + strconv.Itoa(attack.Max) + ")\r\n")
	}

	// Send information message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       monster.Name + " information",
		Description: msg.String(),
	})
	return err
}

func viewMonsterLoot(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, monster *xml.Monster) error {
	msg := strings.Builder{}
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
			msg.WriteString("\r\n- **" + i.Name + "** - " + strconv.FormatFloat(chance, 'f', 2, 64) + "%")
			continue
		}

		msg.WriteString("\r\n- **" + item.Name + "** - " + strconv.FormatFloat(chance, 'f', 2, 64) + "%")
	}

	// Send loot message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       monster.Name + " loot",
		Description: msg.String(),
	})
	return err
}
