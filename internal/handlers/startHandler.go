package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendLanguageSwitcher(messageId int64, bot *tgbotapi.BotAPI) {
	languageKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇦", "uk-UA"),
			tgbotapi.NewInlineKeyboardButtonData("🇬🇧", "en-US"),
		),
	)

	message := tgbotapi.NewMessage(
		messageId,
		"🇬🇧Please select your preferred language\n"+
			"🇺🇦Будь ласка, оберіть бажану мову",
	)

	message.ReplyMarkup = languageKeyboard

	if _, err := bot.Send(message); err != nil {
		panic(err)
	}
}

func handleLanguage(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	println(callbackQuery.Data)
}

func groupWelcome(messageId int64, bot *tgbotapi.BotAPI) {
	sendLanguageSwitcher(messageId, bot)
}

func userWelcome(messageId int64, bot *tgbotapi.BotAPI) {
	sendLanguageSwitcher(messageId, bot)
}

func startHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {

	if !message.Chat.IsGroup() && !message.Chat.IsSuperGroup() {
		userWelcome(message.Chat.ID, bot)
	} else {
		groupWelcome(message.Chat.ID, bot)
	}
}
