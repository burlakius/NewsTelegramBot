package handlers

import (
	"news_telegram_bot/internal/config"
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func adminAuthentication(chatID int64, lang string, bot *tgbotapi.BotAPI) {
	printer := translator.GetPrinter(lang)

	responceMessage := tgbotapi.NewMessage(chatID, printer.Sprintf(
		"Привіт усім!\nЯ - @%v, створений, щоб допомагати адмінам у публікації нових новин та відповідати на питання студентів.\n\nЯ співпрацюю тільки зі своїми адміністраторами, тому пойдіть, будь ласка, аутентифікацію.\n\nВведіть пароль:",
		bot.Self.UserName,
	))

	err := redisdb.SetChatState(chatID, "WaitPassword")
	if err != nil {
		return // TODO!!!!!!!
	}
	bot.Send(responceMessage)
}

func receiveAdminPassword(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	defer redisdb.DoneChatState(message.Chat.ID)

	lang, err := redisdb.GetLanguage(message.Chat.ID)
	if err != nil {
		return // TODO!!!!!!!
	}
	printer := translator.GetPrinter(lang)
	responceMessage := tgbotapi.NewMessage(message.Chat.ID, "")
	if message.Text == config.AdminPassword {
		responceMessage.Text = printer.Sprintf("Пароль вірний ✅")
	} else {
		responceMessage.Text = printer.Sprintf("Пароль невірний ❌")
		defer bot.Send(
			tgbotapi.LeaveChatConfig{ChatID: message.Chat.ID},
		)
	}

	bot.Send(responceMessage)
}
