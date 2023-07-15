package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/text/message"
)

func setMenu(chatID int64, printer *message.Printer, bot *tgbotapi.BotAPI) {

	userKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Новини 📰")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Задати питання ❓")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Змінити мову [🇺🇦|🇬🇧]")),
		),
	)

	responceMessage := tgbotapi.NewMessage(chatID, printer.Sprintf("Меню налаштоване 📱"))
	responceMessage.ReplyMarkup = userKeyboard

	bot.Send(responceMessage)
}
