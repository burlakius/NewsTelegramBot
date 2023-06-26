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
)

var CONFIG config.Config

func init() {
	CONFIG = config.ReadConfig()

	redisdb.RedisConnect(CONFIG.RedisAddress, CONFIG.RedisPort)

	logging.LoggerSetup(CONFIG.LogPath, CONFIG.LogLevel)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(CONFIG.TGbotToken)
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
