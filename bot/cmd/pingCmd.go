package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingRun(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("Running ping command")
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
}
