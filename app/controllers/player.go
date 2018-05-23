package controllers

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

var viewPlayerCommand = Command{
	Usage: "player Name, option",
	Options: []CommandOption{
		{
			Name:        "info",
			Description: "Shows the basic information about a character",
		},
	},
}

// ViewPlayer views the given server character
func ViewPlayer(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	data := strings.Split(m.Content, ",")
	if len(data) <= 1 {
		return viewPlayerCommand.RenderUsage("Unknown option", context, s, m)
	}

	// Retrieve player by name
	player, err := models.GetPlayerByName(context.DB, strings.TrimSpace(data[0]))
	if err != nil {
		return viewPlayerCommand.RenderUsage("Player not found", context, s, m)
	}

	switch strings.TrimSpace(data[1]) {
	case "info":
		return viewPlayerInformation(context, s, m, player)
	default:
		return viewPlayerInformation(context, s, m, player)
	}
}

func viewPlayerInformation(context *Context, s *discordgo.Session, m *discordgo.MessageCreate, player *models.Player) error {
	// Retrieve player vocation
	playerVocation := "No vocation"
	for _, v := range context.Vocations {
		if v.ID == player.Vocation {
			playerVocation = v.Name
			break
		}
	}

	// Create message
	skillsMessage := strings.Builder{}
	skillsMessage.WriteString("- **Magic**: " + strconv.Itoa(player.MagicLevel))
	skillsMessage.WriteString("\r\n- **Fist**: " + strconv.Itoa(player.SkillFist))
	skillsMessage.WriteString("\r\n- **Club**: " + strconv.Itoa(player.SkillClub))
	skillsMessage.WriteString("\r\n- **Sword**: " + strconv.Itoa(player.SkillSword))
	skillsMessage.WriteString("\r\n- **Axe**: " + strconv.Itoa(player.SkillAxe))
	skillsMessage.WriteString("\r\n- **Distance**: " + strconv.Itoa(player.SkillDist))
	skillsMessage.WriteString("\r\n- **Shielding**: " + strconv.Itoa(player.SkillShielding))
	skillsMessage.WriteString("\r\n- **Fishing**: " + strconv.Itoa(player.SkillFishing))
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Vocation",
			Value: playerVocation,
		},
		{
			Name:  "Level",
			Value: "Currently level **" + strconv.Itoa(player.Level) + "** with `" + strconv.FormatInt(player.Experience, 10) + "` experience",
		},
		{
			Name:  "Skills",
			Value: skillsMessage.String(),
		},
	}

	// Send message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  "View player " + player.Name,
		Color:  3447003,
		Fields: fields,
	})
	return err
}
