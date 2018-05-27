package controllers

import (
	"bytes"
	"database/sql"
	"strings"
	"time"

	"text/template"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/xml"
	cache "github.com/robfig/go-cache"
)

// Context main controller for all actions
type Context struct {
	Start     time.Time
	Config    *config.Config
	Monsters  map[string]*xml.Monster
	Vocations map[string]*xml.Vocation
	Items     map[int]xml.Item
	DB        *sql.DB
	Template  *template.Template
	Cache     *cache.Cache
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

// ExecuteTemplate executes the given markdown template file
func (c *Context) ExecuteTemplate(name string, data map[string]interface{}) (string, error) {
	buff := &bytes.Buffer{}
	if err := c.Template.ExecuteTemplate(buff, name+c.Config.Template.Extension, data); err != nil {
		return "", err
	}
	return buff.String(), nil
}

// RenderUsage sends a message with the command usage
func (c *Command) RenderUsage(title string, ctx *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
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
	return &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       title,
		Description: msg.String(),
		Fields:      fields,
	}, nil
}
