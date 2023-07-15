package handlers

import (
	"news_telegram_bot/pkg/databases/mariadb"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getNews(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	newsList, err := mariadb.GetUserNews(message.Chat.ID)
	if err != nil {

		sendBotStorageError(message.Chat.ID, bot)
		panic(err)
	}
	if len(newsList) == 0 {
		noNewsMessage := tgbotapi.NewMessage(message.Chat.ID, printer.Sprintf("Новин немає..."))
		bot.Send(noNewsMessage)
		return
	}

	var news mariadb.News
	for _, news = range newsList {
		dateMessage := tgbotapi.NewMessage(message.Chat.ID, news.PublicationDate)
		bot.Send(dateMessage)

		newsMessage := tgbotapi.NewCopyMessage(message.Chat.ID, news.ChatID, news.MessageID)
		bot.Send(newsMessage)
	}

	err = mariadb.UpdateUserLastNews(news.NewsID, message.Chat.ID)
	if err != nil {
		sendBotStorageError(message.Chat.ID, bot)
	}
}
