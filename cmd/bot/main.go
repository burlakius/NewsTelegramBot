package main

import (
	"context"
	"news_telegram_bot/pkg/config"
	"news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/logging"
	"news_telegram_bot/pkg/router"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	_ "news_telegram_bot/internal/translations"
)

func init() {
	config.LoadConfig()

	redisdb.RedisConnect(config.RedisAddress, config.RedisPort)

	logging.LoggerSetup(config.LogPath, config.LogLevel)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TGbotToken)
	if err != nil {
		logrus.Fatal("BOT ERROR: ", err)
	}

	logrus.Infof("Authorized on account %s", bot.Self.UserName)

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 15

	router := router.NewRouter(bot, update)
	router.InitDispatchers(4)

	ctx, close := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer close()

	router.StartPolling(ctx)
}
