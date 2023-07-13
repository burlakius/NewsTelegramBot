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
		sendBotStorageError(message.Chat.ID, bot)
		return
	}

	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

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
	err := mariadb.AddQuestionMessage(message.From.ID, message.Chat.ID, message.MessageID)
	if err != nil {
		sendBotStorageError(message.Chat.ID, bot)
		return
	}

	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

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
	printer, err := translator.GetPrinterByChatID(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	redisdb.DoneChatState(callbackQuery.Message.Chat.ID)

	responceMessage := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, printer.Sprintf("Питання надіслані\n\nЧекайте, на відповідь!"))
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
	err := mariadb.DeleteQuestionMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.ReplyToMessage.MessageID)
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	printer, err := translator.GetPrinterByChatID(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	editedMessage := tgbotapi.NewEditMessageReplyMarkup(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Видалено..."), "_"),
			),
		),
	)

	bot.Send(editedMessage)
}
