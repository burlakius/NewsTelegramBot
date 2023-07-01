package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"news_telegram_bot/pkg/translator"
)

func setMenu(chatID int64, lang string, bot *tgbotapi.BotAPI) {
	printer := translator.GetPrinter(lang)

	userKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Новини 📰")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Задати питання ❓")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(printer.Sprintf("Мої налаштування ⚙️")),
		),
	)

	responceMessage := tgbotapi.NewMessage(chatID, printer.Sprintf("Меню налаштоване 📱"))
	responceMessage.ReplyMarkup = userKeyboard

	bot.Send(responceMessage)
}
