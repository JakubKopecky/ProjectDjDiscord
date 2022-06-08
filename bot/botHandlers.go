package bot

import (
	"ProjectDjDiscord/bot/cmd"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}

	if !strings.HasPrefix(m.Content, currentConfig.BotPrefix) {
		return
	}

	switch m.Content[1:] {
	case "ping":
		cmd.PingRun(s, m)
	case "voice":
		cmd.VoiceRun(s, m)
	}
}
