package handlers

import (
	redisdb "news_telegram_bot/pkg/databases/redis"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func receiveLanguage(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	defer bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))

	language := callbackQuery.Data
	err := redisdb.SetLanguage(callbackQuery.Message.Chat.ID, language)

	responceTest := map[string]string{
		"uk-UA": "Мова 🇺🇦 налаштована!",
		"en-US": "The language 🇺🇸 is set up!",
	}
	message := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "")
	if err == nil {
		message.Text = responceTest[language]
	} else {
		message.Text = "Oops, an error occurred, try again!\n\nОй, сталася помилка, спробуй ще!"
		defer sendLanguageSwitcher(callbackQuery.Message, bot)

	}
	bot.Send(message)
	startHandler(callbackQuery.Message, bot)
}

func sendLanguageSwitcher(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	languageKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇦", "uk-UA"),
			tgbotapi.NewInlineKeyboardButtonData("🇺🇸", "en-US"),
		),
	)

	responceMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		"🇺🇸Please select your preferred language\n"+
			"🇺🇦Будь ласка, оберіть бажану мову",
	)

	responceMessage.ReplyMarkup = languageKeyboard

	bot.Send(responceMessage)
}
