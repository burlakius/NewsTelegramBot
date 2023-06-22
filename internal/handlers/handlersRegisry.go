package handlers

import (
	"news_telegram_bot/pkg/dispatcher"
	"news_telegram_bot/pkg/filters"
)

func RegisterAllHandlers(dp *dispatcher.Dispatcher) {
	// Message handlers
	dp.RegisterMessageHandler(startHandler, filters.CommandFilter("start"))

	// Edited message handlers

	// Callback query handlers
	dp.RegisterCallbackQueryHandler(handleLanguage, filters.CallbackDataFilter("en-US", "uk-UA"))
}
