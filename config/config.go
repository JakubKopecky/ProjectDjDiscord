package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func LoadConfig() *ConfigStruct {
	log.Print("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var loadedConfig = ConfigStruct{}
	err = json.Unmarshal(file, &loadedConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &loadedConfig
}
