package handlers

import (
	"database/sql"
	"news_telegram_bot/internal/config"
	"news_telegram_bot/pkg/databases/mariadb"
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/text/message"
)

func adminAuthentication(chatID int64, printer *message.Printer, bot *tgbotapi.BotAPI) {
	responceMessage := tgbotapi.NewMessage(chatID, printer.Sprintf(
		"Привіт усім!\nЯ - @%v, створений, щоб допомагати адмінам у публікації нових новин та відповідати на питання студентів.\n\nЯ співпрацюю тільки зі своїми адміністраторами, тому пойдіть, будь ласка, аутентифікацію.\n\nВведіть пароль:",
		bot.Self.UserName,
	))

	if !mariadb.IsAdminChat(chatID) {
		err := redisdb.SetChatState(chatID, "WaitPassword")
		if err != nil {
			sendBotStorageError(chatID, bot)
			return
		}
		bot.Send(responceMessage)
	}
}

func setAdminsCommands(chatID int64, printer *message.Printer, bot *tgbotapi.BotAPI) {
	adminCommands := tgbotapi.NewSetMyCommandsWithScope(
		tgbotapi.NewBotCommandScopeChat(chatID),
		tgbotapi.BotCommand{
			Command:     "get_question",
			Description: printer.Sprintf("Вивести питання користувачів з можливістю відповіді"),
		},
		tgbotapi.BotCommand{
			Command:     "set_news",
			Description: printer.Sprintf("Додати нову звичайну новину"),
		},
		tgbotapi.BotCommand{
			Command:     "set_important_news",
			Description: printer.Sprintf("Додати нову важливу новину"),
		},
		tgbotapi.BotCommand{
			Command:     "edit_news",
			Description: printer.Sprintf("Редагувати або видалити новину"),
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: printer.Sprintf("Детальніше про команди"),
		},
	)

	bot.Send(adminCommands)
}

func receiveAdminPassword(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	defer redisdb.DoneChatState(message.Chat.ID)

	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, "")
	if message.Text == config.AdminPassword {
		responceMessage.Text = printer.Sprintf("Пароль вірний ✅\nНатисніть /help якщо вам потрібно дізнатись, як зі мною співпрацювати, або у вас не відображаються меню команд")
		mariadb.AddNewAdminChat(message.Chat.ID)
		defer setAdminsCommands(message.Chat.ID, printer, bot)
	} else {
		responceMessage.Text = printer.Sprintf("Пароль невірний ❌")
		defer bot.Send(
			tgbotapi.LeaveChatConfig{ChatID: message.Chat.ID},
		)
	}

	bot.Send(responceMessage)
}

func helpForAdmins(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		printer.Sprintf("Я ваш особистий телеграм-бот, готовий допомогти вам з отриманням та управлінням новинами.\n\nОсь команди, які ви можете використовувати:\n\n/get_question - отримати одне питання з можливістю відповіді, або видалення(питання не видалиться до тих пір, поки ви самі цього не захочете)\n/set_news - Додати нову звичайну новину.\n/set_important_news - Додати нову важливу новину.\n\n!Важливі новини відрізняються від звичайних тим, що користувач не зможе відключити собі надходження важливих\n\n/edit_news - Редагувати або видалити новину.\n/help - Детальніше про команди та їх використання.\n\nЯ готовий надати вам всю необхідну інформацію та виконувати ваші запити. Просто надішліть мені одну з цих команд, і я буду радий допомогти!"),
	)

	setAdminsCommands(message.Chat.ID, printer, bot)
	bot.Send(responceMessage)
}

func getQuestion(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, "")
	question, err := mariadb.GetQuestion()
	if err != nil {
		if err == sql.ErrNoRows {
			responceMessage.Text = printer.Sprintf("Питань немає...")
			bot.Send(responceMessage)
			return
		} else {
			sendBotStorageError(message.Chat.ID, bot)
			return
		}
	}

	sendQuestion(
		message.Chat.ID,
		question,
		printer,
		bot,
	)
}

func sendQuestion(chatID int64, question *mariadb.Question, printer *message.Printer, bot *tgbotapi.BotAPI) {
	aboutUserMessage := tgbotapi.NewMessage(chatID, printer.Sprintf("Питання від:\n\n%s\n%s", question.FirstName+" "+question.LastName, question.Username))
	bot.Send(aboutUserMessage)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Відповісти ✍️"), "ReplyToQuestion"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Видалити 🗑"), "DeleteQuestion"),
		),
	)

	questionMessage := tgbotapi.NewCopyMessage(
		chatID,
		question.ChatID,
		question.MessageID,
	)

	questionMessage.ReplyMarkup = inlineKeyboard
	bot.Send(questionMessage)
}

func replyToQuestion(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	redisdb.SetChatState(callbackQuery.Message.Chat.ID, "WaitAnswerMessage")

	printer, err := translator.GetPrinterByChatID(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	changeQuestionMessageMenu(callbackQuery.Message, printer, bot)

	responceMessage := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, printer.Sprintf("Чекаю на вашу відповідь...\n\nУвага! Всі наступні повідомлення будуть зараховані, як відповіть на данне питання і будуть відправленні користувачу"))
	bot.Send(responceMessage)
}

func changeQuestionMessageMenu(message *tgbotapi.Message, printer *message.Printer, bot *tgbotapi.BotAPI) {
	newInlineMenu := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Надіслати відповідь ✅"), "SendAnswer"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Скасувати ❌"), "cancel"),
		),
	)

	newQuestionMessage := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, newInlineMenu)

	bot.Send(newQuestionMessage)
}

func receiveReplyMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	err := mariadb.SaveAnswerMessage(message.Chat.ID, message.MessageID)
	if err != nil {
		sendBotStorageError(message.Chat.ID, bot)
		return
	}

	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, printer.Sprintf("Відповідь збережена..."))
	responceMessage.ReplyToMessageID = message.MessageID
	responceMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(printer.Sprintf("Видалити"), "DeleteAnswerMessage"),
		),
	)
	bot.Send(responceMessage)
}

func deleteAnswerMessage(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	err = mariadb.DeleteAnswerMessage(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.ReplyToMessage.MessageID,
	)
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	editedMessage := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		printer.Sprintf("Видалено..."),
	)
	bot.Send(editedMessage)

	editedMessageMarkup := tgbotapi.NewEditMessageReplyMarkup(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(),
	)
	bot.Send(editedMessageMarkup)
}

func sendAdminAnswer(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	redisdb.DoneChatState(callbackQuery.Message.Chat.ID)

	lang, err := redisdb.GetLanguage(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}
	printer := translator.GetPrinter(lang)

	answerMessages, err := mariadb.GetAllAnswerMessages()
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	question, err := mariadb.GetQuestion()
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	titleMessage := tgbotapi.NewMessage(question.ChatID, printer.Sprintf("Відповідь на питання!"))
	titleMessage.ReplyToMessageID = question.MessageID
	bot.Send(titleMessage)

	for chatID, messageID := range answerMessages {
		answerMessage := tgbotapi.NewCopyMessage(question.UserID, chatID, messageID)
		bot.Send(answerMessage)

		confirmMessage := tgbotapi.NewMessage(chatID, printer.Sprintf("Відповідь надіслана!"))
		confirmMessage.ReplyToMessageID = messageID
		bot.Send(confirmMessage)
	}

	err = mariadb.DeleteAllAnswerMessages()
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	err = mariadb.DeleteQuestionFirstMessage()
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	editedMessageMarkup := tgbotapi.NewEditMessageReplyMarkup(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					printer.Sprintf("Відповідь надісланна..."),
					"_",
				),
			),
		),
	)
	bot.Send(editedMessageMarkup)
}

func deleteQuestion(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	err := mariadb.DeleteQuestionFirstMessage()
	if err != nil {
		sendBotStorageError(callbackQuery.Message.Chat.ID, bot)
		return
	}

	lang, err := redisdb.GetLanguage(callbackQuery.Message.Chat.ID)
	if err != nil {
		sendLanguageError(callbackQuery.Message.Chat.ID, bot)
		return
	}
	printer := translator.GetPrinter(lang)

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
