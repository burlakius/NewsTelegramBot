package handlers

import (
	"news_telegram_bot/pkg/databases/mariadb"
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func userQuestionHandler(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	err := redisdb.SetChatState(message.Chat.ID, "WaitQuestion")
	if err != nil {
		return //TODO!!!!!
	}

	lang, err := redisdb.GetLanguage(message.Chat.ID)
	if err != nil {
		return //TODO!!!!!
	}
	printer := translator.GetPrinter(lang)

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, printer.Sprintf("Очікую на ваше питання\n\nУВАГА! Всі наступні повідомлення будуть враховуватись, як питання і будуть надісланні адміністраторам"))
	responceMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprint("Надіслати питання ✉️"), "SendQuestions"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprint("Відмінити ❌"), "cancel"),
		),
	)

	bot.Send(responceMessage)
}

func receiveQuetionMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	err := mariadb.SetQuestion(message.Chat.ID, message.MessageID)
	if err != nil {
		panic(err) // TODO!!!!!!!
	}

	lang, err := redisdb.GetLanguage(message.Chat.ID)
	if err != nil {
		return //TODO!!!!!
	}
	printer := translator.GetPrinter(lang)
	responceMessage := tgbotapi.NewMessage(message.Chat.ID, printer.Sprintf("Питання збережено..."))
	responceMessage.ReplyToMessageID = message.MessageID
	responceMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				printer.Sprintf("Видалити повідомлення"),
				"DeleteQuestionMessage",
			),
		),
	)

	bot.Send(responceMessage)
}

func sendUserQuestions(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	lang, err := redisdb.GetLanguage(callbackQuery.Message.Chat.ID)
	if err != nil {
		return //TODO!!!!!
	}
	printer := translator.GetPrinter(lang)

	redisdb.DoneChatState(callbackQuery.Message.Chat.ID)

	responceMessage := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, printer.Sprintf("Питання надіслані\n\nЧекайте, на відповідь"))
	bot.Send(responceMessage)

	editedMessageMarkup := tgbotapi.NewEditMessageReplyMarkup(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					printer.Sprintf("Питання надісланне..."),
					"_",
				),
			),
		),
	)
	bot.Send(editedMessageMarkup)
}

func deleteUserQuestion(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	err := redisdb.DoneChatState(callbackQuery.Message.Chat.ID)
	if err != nil {
		panic(err) // TODO!!!!!!!!!!!
	}

	err = mariadb.DeleteQuestionMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.ReplyToMessage.MessageID)
	if err != nil {
		panic(err) //TODO!!!!!!!!!!
	}

	lang, err := redisdb.GetLanguage(callbackQuery.Message.Chat.ID)
	if err != nil {
		return //TODO!!!!!
	}
	printer := translator.GetPrinter(lang)

	responceMessage := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, printer.Sprintf("Повідомлення видалено..."))
	responceMessage.ReplyToMessageID = callbackQuery.Message.ReplyToMessage.MessageID

	bot.Send(responceMessage)
}
