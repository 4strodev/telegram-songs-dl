package main

import (
	"log"

	"github.com/4strodev/songs_dl_telegram/internal"
	"github.com/4strodev/songs_dl_telegram/internal/environment"
	"github.com/4strodev/songs_dl_telegram/internal/songs"
)

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN"`
	SongsFile     string `env:"SONGS_FILE"`
}

func main() {
	config := Config{}
	err := environment.LoadEnvironmentVariables(&config)
	if err != nil {
		log.Fatal(err)
	}

	bot := internal.Bot{
		Token: config.TelegramToken,
		SongsRepository: songs.SongsRepository{
			Destination: config.SongsFile,
		},
	}

	bot.Start()
}
