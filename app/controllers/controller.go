package controllers

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/xml"
)

// Context main controller for all actions
type Context struct {
	Start     time.Time
	Config    *config.Config
	Monsters  map[string]*xml.Monster
	Vocations map[string]*xml.Vocation
	Items     map[int]xml.Item
	DB        *sql.DB
}

// Command defines a discord command
type Command struct {
	Usage       string
	Description string
	Options     []CommandOption
}

// CommandOption defines a command option value
type CommandOption struct {
	Name        string
	Usage       string
	Description string
}

// RenderUsage sends a message with the command usage
func (c *Command) RenderUsage(title string, ctx *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	msg := strings.Builder{}
	msg.WriteString(c.Description)
	msg.WriteString("\r\nCommand usage: `" + ctx.Config.Discord.Prefix + c.Usage + "`\r\n")
	fields := []*discordgo.MessageEmbedField{}

	// Add usage options if needed
	if len(c.Options) > 0 {
		for _, o := range c.Options {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  o.Name,
				Value: o.Description,
			})
		}
	}

	// Send usage message
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       title,
		Description: msg.String(),
		Fields:      fields,
	})
	return err
}

func timeAgo(a time.Time, b time.Time) string {
	y, m, d, h, x, s := diff(a, b)
	msg := ""

	// Render message as year
	if y > 0 {
		msg += strconv.Itoa(y)
		if y == 1 {
			msg += " year"
		} else {
			msg += " years"
		}
		return msg
	}

	// Render message as month
	if m > 0 {
		msg += strconv.Itoa(m)
		if m == 1 {
			msg += " month"
		} else {
			msg += " months"
		}
		return msg
	}

	// Render message as day
	if d > 0 {
		msg += strconv.Itoa(d)
		if d == 1 {
			msg += " day"
		} else {
			msg += " days"
		}
		return msg
	}

	// Render message as hour
	if h > 0 {
		msg += strconv.Itoa(h)
		if h == 1 {
			msg += " hour"
		} else {
			msg += " hours"
		}
		return msg
	}

	// Render message as minute
	if x > 0 {
		msg += strconv.Itoa(x)
		if x == 1 {
			msg += " minute"
		} else {
			msg += " minutes"
		}
		return msg
	}

	// Render message as second
	if s > 0 {
		msg += strconv.Itoa(s)
		if s == 1 {
			msg += " second"
		} else {
			msg += " seconds"
		}
		return msg
	}

	return ""
}
