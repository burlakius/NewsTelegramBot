package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendContacts(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	responceMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		"Coming soon..", //TODO!
	)

	bot.Send(responceMessage)
}
