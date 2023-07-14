package handlers

import (
	"news_telegram_bot/pkg/databases/mariadb"
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
		err := mariadb.AddNewUser(
			message.Chat.ID,
			message.From.FirstName,
			message.From.LastName,
			message.From.UserName,
		)
		if err != nil {
			sendNewUserError(message.From.ID, bot)
			return
		}

		setMenu(message.Chat.ID, printer, bot)
	} else if message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
		adminAuthentication(message.Chat.ID, printer, bot)
	}
}
