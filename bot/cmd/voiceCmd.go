package cmd

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func VoiceRun(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("Running voice command")
	voiceChannel, err := s.State.VoiceState(m.GuildID, m.Author.ID)
	if err != nil {
		log.Print("voice channel not found")
		s.ChannelMessageSend(m.ChannelID, "Voice channel not found")
		return
	}

	vc, err := s.ChannelVoiceJoin(m.GuildID, voiceChannel.ChannelID, false, false)
	if err != nil {
		log.Panic(err.Error())
	}

	time.Sleep(250 * time.Millisecond)
	vc.Speaking(true)
	time.Sleep(1000 * time.Millisecond)
	vc.Speaking(false)
	time.Sleep(250 * time.Millisecond)
	vc.Disconnect()
}
