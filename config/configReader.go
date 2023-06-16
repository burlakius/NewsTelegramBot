package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type config struct {
	TGbotToken string
	LogLevel   logrus.Level
	LogPath    string
}

func ReadConfig(pathToFile string) config {
	godotenv.Load(pathToFile)

	logLevels := map[string]logrus.Level{
		"debug":   logrus.DebugLevel,
		"info":    logrus.InfoLevel,
		"warning": logrus.WarnLevel,
		"error":   logrus.ErrorLevel,
		"fatal":   logrus.FatalLevel,
	}

	return config{
		TGbotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		LogLevel:   logLevels[os.Getenv("LOGGING_LEVEL")],
		LogPath:    os.Getenv("PATH_TO_LOG_FILE"),
	}
}
