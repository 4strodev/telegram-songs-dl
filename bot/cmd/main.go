package main

import (
	"os"

	"github.com/4strodev/songs_dl_telegram/internal/songs"
	"github.com/4strodev/songs_dl_telegram/internal"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	songsFile := os.Getenv("SONGS_FILE")
	bot := internal.Bot{
		Token: token,
		SongsRepository: songs.SongsRepository{
			Destination: songsFile,
		},
	}

	bot.Start()
}
