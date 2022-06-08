package bot

import (
	"ProjectDjDiscord/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	currentConfig *config.ConfigStruct
)

type botStruct struct {
	botSession *discordgo.Session
}

func NewBot(config *config.ConfigStruct) (*botStruct, error) {
	currentConfig = config
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
