package main

import (
	"fmt"

	"ProjectDjDiscord/bot"
	"ProjectDjDiscord/config"
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
