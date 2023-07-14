package handlers

import (
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	"news_telegram_bot/pkg/databases/mariadb"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/text/message"
)

func addNews(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	switch message.Command() {
	case "set_news":
		redisdb.SetChatState(message.Chat.ID, "WaitNews")
	case "set_important_news":
		redisdb.SetChatState(message.Chat.ID, "WaitImportantNews")
	default:
		return
	}
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		printer.Sprintf("Чекаю на новини!\n\nУвага! Всі наступні повідомлення будуть зараховані, як новини і будуть відправленні користувачу"),
	)
	responceMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Надіслати новини 📰"), "DoneAddingNews"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Відмінити ❌"), "cancel"),
		),
	)
	bot.Send(responceMessage)
}

func receiveNewsMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	chatState, err := redisdb.GetChatState(message.Chat.ID)
	if err != nil {
		sendChatStateError(message.Chat.ID, bot)
		return
	}
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	var newsType string
	switch chatState {
	case "WaitNews":
		newsType = "regular"
	case "WaitImportantNews":
		newsType = "important"
	}

	err = mariadb.AddNewsMessage(
		message.Chat.ID,
		message.MessageID,
		newsType,
	)
	if err != nil {
		sendBotStorageError(message.Chat.ID, bot)
		redisdb.DoneChatState(message.Chat.ID)
	}

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, printer.Sprintf("Новина збережена"))
	responceMessage.ReplyToMessageID = message.MessageID
	responceMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Видалити 🗑"), "DeleteNewsMessage"),
		),
	)
	bot.Send(responceMessage)
}

func deleteNewsMessage(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	err := mariadb.DeleteNewsMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.ReplyToMessage.MessageID)
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

func doneAddingNews(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	defer redisdb.DoneChatState(callbackQuery.Message.Chat.ID)
	chatState, err := redisdb.GetChatState(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendChatStateError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	printer, err := translator.GetPrinterByChatID(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}
	responceMessage := tgbotapi.NewMessage(
		callbackQuery.Message.Chat.ID,
		printer.Sprintf("Додання новин завершено!"),
	)
	bot.Send(responceMessage)

	editedMessage := tgbotapi.NewEditMessageReplyMarkup(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Новини надіслані..."), "_"),
			),
		),
	)
	bot.Send(editedMessage)

	switch chatState {
	case "WaitNews":
		sendNews(callbackQuery.Message.Chat.ID, "regular", printer, bot)
	case "WaitImportantNews":
		sendNews(callbackQuery.Message.Chat.ID, "important", printer, bot)
	}
}

func sendNews(adminChatID int64, newsType string, printer *message.Printer, bot *tgbotapi.BotAPI) {
	newsList, err := mariadb.GetAllHiddenNews(newsType)
	if err != nil {
		sendBotStorageError(adminChatID, bot)
		return
	}
	targetUsers, err := mariadb.GetTargetUsers(newsType)
	if err != nil {
		sendBotStorageError(adminChatID, bot)
		return
	}

	for _, news := range newsList {
		for _, userID := range targetUsers {
			newsMessage := tgbotapi.NewCopyMessage(userID, news.ChatID, news.MessageID)
			bot.Send(newsMessage)
		}
	}

	err = mariadb.UnhideNews()
	if err != nil {
		sendBotStorageError(adminChatID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(adminChatID, printer.Sprintf("Новини надіслані 🥳"))
	bot.Send(responceMessage)
}
