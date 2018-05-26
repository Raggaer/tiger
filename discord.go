package main

import "github.com/bwmarrin/discordgo"

func getChannelName(s *discordgo.Session, channel string) (string, error) {
	c, err := s.Channel(channel)
	if err != nil {
		return "", err
	}
	return c.Name, nil
}
