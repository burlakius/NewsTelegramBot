package handlers

import (
	"news_telegram_bot/pkg/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"news_telegram_bot/internal/databases/redis"
)

func startHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI, botState *state.BotState) {
	chatLocale, err := redisdb.GetLanguage(message.Chat.ID)
	if err != nil {
		sendLanguageSwitcher(message.Chat.ID, bot)
		return
	}

	if message.Chat.IsPrivate() {
		setMenu(message.Chat.ID, chatLocale, bot)
	} else if message.Chat.IsGroup() {
		println("Chat is GROUP. TODO !")
	}
}
