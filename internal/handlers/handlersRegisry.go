package handlers

import (
	"news_telegram_bot/pkg/dispatcher"
	"news_telegram_bot/pkg/filters"
)

func RegisterAllHandlers(dp *dispatcher.Dispatcher) {
	// Message handlers
	dp.RegisterMessageHandler(startHandler, filters.CommandFilter([]string{"start"}, "*"))

	// Edited message handlers

	// Callback query handlers
	dp.RegisterCallbackQueryHandler(receiveLanguage, filters.CallbackDataFilter([]string{"en-US", "uk-UA"}, "*"))
}
