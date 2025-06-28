package main

import (
	"log"

	"github.com/4strodev/songs_dl_telegram/internal"
	"github.com/4strodev/songs_dl_telegram/internal/environment"
	"github.com/4strodev/songs_dl_telegram/internal/songs"
	"github.com/4strodev/songs_dl_telegram/services"
)

type Config struct {
	TelegramToken    string `env:"TELEGRAM_TOKEN"`
	SongsFile        string `env:"SONGS_FILE"`
	SongsDirectory   string `env:"SONGS_DIR"`
	ScriptsDirectory string `env:"SCRIPTS_DIR"`
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
		Downloader: services.Downloader{
			ScriptsFolder:  config.ScriptsDirectory,
			SongsDirectory: config.SongsDirectory,
		},
	}

	bot.Start()
}
