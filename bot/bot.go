package bot

import (
	"ProjectDjDiscord/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

type botStruct struct {
	botSession *discordgo.Session
}

func NewBot(config *config.ConfigStruct) (*botStruct, error) {
	botSession, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	botSession.AddHandler(messageHandler)

	return &botStruct{botSession: botSession}, nil
}

func (bot *botStruct) Start() error {
	err := bot.botSession.Open()
	if err != nil {
		return err
	}

	log.Print("Bot is running!")

	return nil
}

func (bot *botStruct) Close() {
	log.Print("Closing bot...")
	bot.botSession.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
