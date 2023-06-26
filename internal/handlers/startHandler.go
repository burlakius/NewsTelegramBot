package handlers

import (
	"news_telegram_bot/pkg/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func groupWelcome(chatID int64, bot *tgbotapi.BotAPI) {
	sendLanguageSwitcher(chatID, bot)
}

func userWelcome(chatID int64, bot *tgbotapi.BotAPI) {
	sendLanguageSwitcher(chatID, bot)
}

func startHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI, botState *state.BotState) {
	if !message.Chat.IsGroup() && !message.Chat.IsSuperGroup() {
		userWelcome(message.Chat.ID, bot)
	} else {
		groupWelcome(message.Chat.ID, bot)
	}
}
