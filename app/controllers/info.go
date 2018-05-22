package controllers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Version returns the current running version
func Version(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Tiger version",
		Description: "Tiger is currently running version **BETA**",
	})
	return err
}

// Uptime returns the current bot uptime
func Uptime(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Retrieve time diff
	_, _, days, hours, minutes, seconds := diff(context.Start, time.Now())

	// Retrieve uptime message
	msg := uptimeMessage(days, hours, minutes, seconds)

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 3447003,
		Title: "Tiger uptime",
		Description: fmt.Sprintf(
			"Tiger has been running for **"+msg+"**",
			days,
			hours,
			minutes,
			seconds,
		),
	})
	return err
}

func uptimeMessage(day, hour, min, sec int) string {
	msg := "%d day"
	if day == 0 || day > 1 {
		msg += "s"
	}
	msg += ", %d hour"
	if hour == 0 || hour > 1 {
		msg += "s"
	}
	msg += ", %d minute"
	if min == 0 || min > 1 {
		msg += "s"
	}
	msg += ", %d second"
	if sec == 0 || sec > 1 {
		msg += "s"
	}
	return msg
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
