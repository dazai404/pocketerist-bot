package telegram

import (
	"log"

	"github.com/dazai404/pocketerist-bot/pkg/config"
	"github.com/dazai404/pocketerist-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string
	messages        config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, tokenRepository repository.TokenRepository, redirectURL string, messages config.Messages) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, tokenRepository: tokenRepository, redirectURL: redirectURL, messages: messages}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
