package filters

import (
	"strings"

	redisdb "news_telegram_bot/pkg/databases/redis"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CommandFilter(commands ...string) func(*tgbotapi.Message) bool {
	return func(message *tgbotapi.Message) bool {
		for _, command := range commands {
			if strings.TrimSpace(message.Text) == "/"+command {
				return true
			}
		}
		return false
	}
}

func MessageTextFilter(texts ...string) func(*tgbotapi.Message) bool {
	return func(message *tgbotapi.Message) bool {
		for _, text := range texts {
			if message.Text == text {
				return true
			}
		}

		return false
	}
}

func CallbackDataFilter(data ...string) func(*tgbotapi.CallbackQuery) bool {
	return func(callbackQuery *tgbotapi.CallbackQuery) bool {
		for _, d := range data {
			if callbackQuery.Data == d {
				return true
			}
		}

		return false
	}
}

func StateFilter(states ...string) func(*tgbotapi.Message) bool {
	return func(message *tgbotapi.Message) bool {
		chatState, err := redisdb.GetChatState(message.Chat.ID)
		if err != nil {
			chatState = "*"
		}
		for _, s := range states {
			if chatState == s {
				return true
			}
		}

		return false
	}
}
