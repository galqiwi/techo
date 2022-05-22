package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type EchoBot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(token string) *EchoBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		err = fmt.Errorf("bot initialization error: %v", err)
		log.Panic(err)
	}
	return &EchoBot{
		bot: bot,
	}
}

func (b *EchoBot) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
