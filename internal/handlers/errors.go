package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendError(chatID int64, bot *tgbotapi.BotAPI, errorText string) {
	message := tgbotapi.NewMessage(chatID, errorText)

	bot.Send(message)
}
