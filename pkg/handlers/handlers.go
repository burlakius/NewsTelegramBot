package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler interface {
	CallbackFunc()
	Check()
}

type MessageHandler struct {
	CallbackFunc func(*tgbotapi.Message, *tgbotapi.BotAPI)
	Filters      []func(*tgbotapi.Message) bool
}

func (mh *MessageHandler) Check(message *tgbotapi.Message) bool {
	for _, filter := range mh.Filters {
		if !filter(message) {
			return false
		}
	}

	return true
}

type CallbackQueryHandler struct {
	CallbackFunc func(*tgbotapi.CallbackQuery, *tgbotapi.BotAPI)
	Filters      []func(*tgbotapi.CallbackQuery) bool
}

func (cqh *CallbackQueryHandler) Check(callbackQuery *tgbotapi.CallbackQuery) bool {
	for _, filter := range cqh.Filters {
		if !filter(callbackQuery) {
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
