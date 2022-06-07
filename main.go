package main

import (
	"fmt"

	"github.com/JakubKopecky/ProjectDjDiscord/bot"
	"github.com/JakubKopecky/ProjectDjDiscord/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
}
