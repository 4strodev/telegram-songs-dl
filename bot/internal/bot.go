package internal

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/4strodev/songs_dl_telegram/internal/songs"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	Token           string
	SongsRepository songs.SongsRepository
	LogHandler      slog.Handler
	logger          *slog.Logger
}

func (b *Bot) Start() {
	if b.LogHandler == nil {
		b.LogHandler = slog.NewJSONHandler(os.Stdout, nil)
	}
	b.logger = slog.New(b.LogHandler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(b.HandleMessage),
	}

	tBot, err := bot.New(b.Token, opts...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting bot")
	tBot.Start(ctx)
}

func (b *Bot) HandleMessage(ctx context.Context, tBot *bot.Bot, update *models.Update) {
	url := strings.Trim(update.Message.Text, "\n\t")
	if url == "" {
		return
	}

	err := b.SongsRepository.AddSong(url)
	if err != nil {
		message := err.Error()
		b.logger.Error("error saving url {error}", "error", err)
		// Send message
		_, err := tBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   message,
		})
		if err != nil {
			b.logger.Error("error sending message to chat {error}", "error", err)
		}
		return
	}

	_, err = tBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "URL saved successfully!",
	})
	if err != nil {
		b.logger.Error("error sending message to chat {error}", "error", err)
	}

	b.logger.Info("new song saved {url}", "url", url)
}
