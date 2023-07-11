package handlers

import (
	redisdb "news_telegram_bot/pkg/databases/redis"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	chatLocale, err := redisdb.GetLanguage(message.Chat.ID)
	if err != nil {
		sendLanguageSwitcher(message.Chat.ID, bot)
		return
	}

	if message.Chat.IsPrivate() {
		setMenu(message.Chat.ID, chatLocale, bot)
	} else if message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
		adminAuthentication(message.Chat.ID, chatLocale, bot)
	}
}
