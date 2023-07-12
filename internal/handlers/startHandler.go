package handlers

import (
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageSwitcher(message, bot)
		return
	}

	if message.Chat.IsPrivate() {
		setMenu(message.Chat.ID, printer, bot)
	} else if message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
		adminAuthentication(message.Chat.ID, printer, bot)
	}
}
