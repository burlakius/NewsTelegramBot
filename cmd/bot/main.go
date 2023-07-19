package main

import (
	"context"
	"news_telegram_bot/internal/config"
	"news_telegram_bot/pkg/databases/mariadb"
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/logging"
	"news_telegram_bot/pkg/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "news_telegram_bot/internal/translations"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func init() {
	time.Sleep(10 * time.Second)

	config.LoadConfig()

	logging.LoggerSetup(config.LogPath, config.LogLevel)

	mariadb.MariadbConnect(config.MariaDBUser, config.MariaDBPassword, config.MariaDBHost, config.MariaDBDatabase)
	redisdb.RedisConnect(config.RedisLanguageSessionsHost, config.RedisChatStatesHost)

	translator.SetupTranslations()
}

func main() {
	defer mariadb.MariadbClose()
	bot, err := tgbotapi.NewBotAPI(config.TGbotToken)
	if err != nil {
		logrus.Fatal(err)
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
