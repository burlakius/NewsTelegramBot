package handlers

import (
	"news_telegram_bot/pkg/dispatcher"
	"news_telegram_bot/pkg/filters"
	"news_telegram_bot/pkg/translator"
)

func RegisterAllHandlers(dp *dispatcher.Dispatcher) {
	// Message handlers
	dp.RegisterMessageHandler(startHandler, filters.CommandFilter("start"))
	dp.RegisterMessageHandler(receiveAdminPassword, filters.StateFilter("WaitPassword"))
	dp.RegisterMessageHandler(helpForAdmins, filters.CommandFilter("help"), filters.AdminChatFilter())
	dp.RegisterMessageHandler(getQuestion, filters.CommandFilter("get_question"), filters.AdminChatFilter())

	dp.RegisterMessageHandler(userQuestionHandler, filters.MessageTextFilter(translator.GetAllTranslations("Задати питання ❓")...))
	dp.RegisterMessageHandler(receiveQuetionMessage, filters.StateFilter("WaitQuestion"))

	dp.RegisterMessageHandler(receiveReplyMessage, filters.StateFilter("WaitAnswerMessage"))

	// Edited message handlers

	// Callback query handlers
	dp.RegisterCallbackQueryHandler(receiveLanguage, filters.CallbackDataFilter("en-US", "uk-UA"))

	dp.RegisterCallbackQueryHandler(sendUserQuestions, filters.CallbackDataFilter("SendQuestions"))
	dp.RegisterCallbackQueryHandler(deleteUserQuestion, filters.CallbackDataFilter("DeleteQuestionMessage"))

	dp.RegisterCallbackQueryHandler(deleteQuestion, filters.CallbackDataFilter("DeleteQuestion"))
	dp.RegisterCallbackQueryHandler(replyToQuestion, filters.CallbackDataFilter("ReplyToQuestion"))
	dp.RegisterCallbackQueryHandler(sendAdminAnswer, filters.CallbackDataFilter("SendAnswer"))
	dp.RegisterCallbackQueryHandler(deleteAnswerMessage, filters.CallbackDataFilter("DeleteAnswerMessage"))

	dp.RegisterCallbackQueryHandler(cancel, filters.CallbackDataFilter("cancel"))
}
