package handlers

import (
	"news_telegram_bot/pkg/dispatcher"
	"news_telegram_bot/pkg/filters"
	"news_telegram_bot/pkg/translator"
)

func RegisterAllHandlers(dp *dispatcher.Dispatcher) {
	// Message handlers
	// 	Common handlers
	dp.RegisterMessageHandler(startHandler, filters.CommandFilter("start"))
	dp.RegisterMessageHandler(sendLanguageSwitcher, filters.CommandFilter("language"))
	dp.RegisterMessageHandler(sendContacts, filters.CommandFilter("contacts"))

	//	Admins handlers
	dp.RegisterMessageHandler(receiveAdminPassword, filters.StateFilter("WaitPassword"))
	dp.RegisterMessageHandler(helpForAdmins, filters.CommandFilter("help"), filters.AdminChatFilter())
	dp.RegisterMessageHandler(getQuestion, filters.CommandFilter("get_question"), filters.AdminChatFilter())
	dp.RegisterMessageHandler(receiveReplyMessage, filters.StateFilter("WaitAnswerMessage"))
	dp.RegisterMessageHandler(addNews, filters.CommandFilter("set_news", "set_important_news"), filters.AdminChatFilter())
	dp.RegisterMessageHandler(receiveNewsMessage, filters.StateFilter("WaitNews", "WaitImportantNews"))

	//	Users handlers
	dp.RegisterMessageHandler(userQuestionHandler, filters.MessageTextFilter(translator.GetAllTranslations("Задати питання ❓")...))
	dp.RegisterMessageHandler(receiveQuetionMessage, filters.StateFilter("WaitQuestion"))

	// Edited message handlers

	// Callback query handlers
	// 	Common handlers
	dp.RegisterCallbackQueryHandler(receiveLanguage, filters.CallbackDataFilter("en-US", "uk-UA"))
	dp.RegisterCallbackQueryHandler(cancel, filters.CallbackDataFilter("cancel"))

	//	Admins handlers
	dp.RegisterCallbackQueryHandler(answetToQuestion, filters.CallbackDataFilter("AnswerToQuestion"))
	dp.RegisterCallbackQueryHandler(deleteQuestion, filters.CallbackDataFilter("DeleteQuestion"))
	dp.RegisterCallbackQueryHandler(deleteAnswerMessage, filters.CallbackDataFilter("DeleteAnswerMessage"))
	dp.RegisterCallbackQueryHandler(sendAdminAnswer, filters.CallbackDataFilter("SendAnswer"))
	dp.RegisterCallbackQueryHandler(doneAddingNews, filters.CallbackDataFilter("DoneAddingNews"))
	dp.RegisterCallbackQueryHandler(deleteNewsMessage, filters.CallbackDataFilter("DeleteNewsMessage"))

	//	Users handlers
	dp.RegisterCallbackQueryHandler(sendUserQuestions, filters.CallbackDataFilter("SendQuestions"))
	dp.RegisterCallbackQueryHandler(deleteUserQuestion, filters.CallbackDataFilter("DeleteQuestionMessage"))

}
