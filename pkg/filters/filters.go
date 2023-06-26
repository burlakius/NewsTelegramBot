package filters

import (
	"strings"

	"news_telegram_bot/pkg/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CommandFilter(commands []string, handlerState string) func(*tgbotapi.Message, *state.BotState) bool {
	return func(message *tgbotapi.Message, botState *state.BotState) bool {
		if botState.GetState() != handlerState {
			return false
		}
		for _, command := range commands {
			if strings.TrimSpace(message.Text) == "/"+command {
				return true
			}
		}
		return false
	}
}

func MessageTextFilter(texts []string, handlerState string) func(*tgbotapi.Message, *state.BotState) bool {
	return func(message *tgbotapi.Message, botState *state.BotState) bool {
		if botState.GetState() != handlerState {
			return false
		}
		for _, text := range texts {
			if message.Text == text {
				return true
			}
		}

		return false
	}
}

func CallbackDataFilter(data []string, handlerState string) func(*tgbotapi.CallbackQuery, *state.BotState) bool {
	return func(callbackQuery *tgbotapi.CallbackQuery, botState *state.BotState) bool {
		if botState.GetState() != handlerState {
			return false
		}
		for _, d := range data {
			if callbackQuery.Data == d {
				return true
			}
		}

		return false
	}
}
