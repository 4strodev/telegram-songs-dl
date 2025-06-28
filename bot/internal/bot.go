package internal

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/4strodev/songs_dl_telegram/internal/songs"
	"github.com/4strodev/songs_dl_telegram/services"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	Token           string
	SongsRepository songs.SongsRepository
	LogHandler      slog.Handler
	logger          *slog.Logger
	Downloader      services.Downloader
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

	b.logger.Info("Starting bot")
	tBot.Start(ctx)
}

func (b *Bot) HandleMessage(ctx context.Context, tBot *bot.Bot, update *models.Update) {
	url := strings.Trim(update.Message.Text, "\n\t")
	if url == "" {
		return
	}

	err := b.SongsRepository.AddSong(url)
	if err != nil {
		b.logger.Error("error saving url {error}", "error", err)
		b.writeError(err, ctx, tBot, update)
		return
	}

	err = b.Downloader.DownloadSong(url)
	if err != nil {
		b.logger.Error("error downloading songs {error}", "error", err)
		b.writeError(err, ctx, tBot, update)
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

func (b *Bot) writeError(err error, ctx context.Context, tBot *bot.Bot, update *models.Update) {
	message := err.Error()
	// Send message
	_, err = tBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   message,
	})
	if err != nil {
		b.logger.Error("error sending message to chat {error}", "error", err)
	}
}
