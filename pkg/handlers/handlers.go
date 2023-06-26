package handler

import (
	"news_telegram_bot/pkg/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler interface {
	CallbackFunc()
	Check()
}

type MessageHandler struct {
	CallbackFunc func(*tgbotapi.Message, *tgbotapi.BotAPI, *state.BotState)
	Filters      []func(*tgbotapi.Message, *state.BotState) bool
}

func (mh *MessageHandler) Check(message *tgbotapi.Message, botState *state.BotState) bool {
	for _, filter := range mh.Filters {
		if !filter(message, botState) {
			return false
		}
	}

	return true
}

type CallbackQueryHandler struct {
	CallbackFunc func(*tgbotapi.CallbackQuery, *tgbotapi.BotAPI, *state.BotState)
	Filters      []func(*tgbotapi.CallbackQuery, *state.BotState) bool
}

func (cqh *CallbackQueryHandler) Check(callbackQuery *tgbotapi.CallbackQuery, botState *state.BotState) bool {
	for _, filter := range cqh.Filters {
		if !filter(callbackQuery, botState) {
			return false
		}
	}

	return true
}

// InlineQueryHandler
// ChosenInlineResultHandler

// ShippingQueryHandler
// PreCheckoutQueryHandler

// PollHandler
// PollAnswerHandler

// MyChatMemberHandler
// ChatMemberHandler
// ChatJoinRequestHandler
