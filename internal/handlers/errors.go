package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendError(chatID int64, bot *tgbotapi.BotAPI, errorText string) {
	message := tgbotapi.NewMessage(chatID, errorText)

	bot.Send(message)
}

func sendLanguageError(chatID int64, bot *tgbotapi.BotAPI) {
	errorMessage := tgbotapi.NewMessage(
		chatID,
		"Помилка!\n\nВибачте, але я не можу зрозуміти, якою мовою відповідати вам. Для вирішення цієї проблеми, будь ласка, спробуйте налаштувати мову, використовуючи команду /language. Якщо це не спрацювало, спробуйте ще раз, або зв'яжіться з адміністраторами для отримання допомоги за допомогою команди /contacts.\n\n\n"+
			"Error!\n\nI'm sorry, but I can't determine the language to respond to you. To resolve this issue, please try setting your language using the command /language. If that didn't work, please try again later, or contact the administrators for assistance using the command /contacts.",
	)
	bot.Send(errorMessage)
}

func sendChatStateError(chatID int64, bot *tgbotapi.BotAPI) {
	errorMessage := tgbotapi.NewMessage(
		chatID,
		"Помилка!\n\nВибачте, але я не можу виконати цю операцію, оскільки не маю доступу до сховища сессій. Для вирішення цієї проблеми, будь ласка, зверніться до адміністраторів за допомогою команди /contacts і повідомте їм про дану помилку.\n\n\n"+
			"Error!\n\nI'm sorry, but I cannot perform this operation as I don't have access to the sessions storage. To resolve this issue, please contact the administrators using the command /contacts and inform them about this error.",
	)

	bot.Send(errorMessage)
}

func sendBotStorageError(chatID int64, bot *tgbotapi.BotAPI) {
	errorMessage := tgbotapi.NewMessage(
		chatID,
		"Помилка!\n\nВибачте, але я не можу виконати цю операцію, оскільки не маю доступу до сховища даних. Для вирішення цієї проблеми, будь ласка, зверніться до адміністраторів за допомогою команди /contacts і повідомте їм про дану помилку.\n\n\n"+
			"Error!\n\nI'm sorry, but I cannot perform this operation as I don't have access to the storage. To resolve this issue, please contact the administrators using the command /contacts and inform them about this error.",
	)

	bot.Send(errorMessage)
}

func sendNewUserError(chatID int64, bot *tgbotapi.BotAPI) {
	errorMessage := tgbotapi.NewMessage(
		chatID,
		"Помилка!\n\nВибачте, але я не можу зберегти дані про вас, оскільки не маю доступу до сховища даних. Для вирішення цієї проблеми, будь ласка, зверніться до адміністраторів за допомогою команди /contacts і повідомте їм про дану помилку."+
			"Error!\n\nI'm sorry, but I cannot save your data as I don't have access to the storage. To resolve this issue, please contact the administrators using the command /contacts and inform them about this error.",
	)

	bot.Send(errorMessage)
}
