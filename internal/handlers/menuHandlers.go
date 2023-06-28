package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func setMenu(chatID int64, lang string, bot *tgbotapi.BotAPI) {
	languageTag := language.MustParse(lang)
	fmt.Println(lang, languageTag)
	printer := message.NewPrinter(languageTag)

	messageText := printer.Sprintf("Меню")
	bot.Send(tgbotapi.NewMessage(chatID, messageText))
}
