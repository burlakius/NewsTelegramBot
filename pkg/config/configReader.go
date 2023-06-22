package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	TGbotToken string
	LogLevel   logrus.Level
	LogPath    string
}

func ReadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalln("Loading .env file error")
	}

	logLevels := map[string]logrus.Level{
		"debug":   logrus.DebugLevel,
		"info":    logrus.InfoLevel,
		"warning": logrus.WarnLevel,
		"error":   logrus.ErrorLevel,
		"fatal":   logrus.FatalLevel,
	}

	level := os.Getenv("LOGGING_LEVEL")
	if level == "" {
		level = "debug"
	}

	return Config{
		TGbotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		LogLevel:   logLevels[level],
		LogPath:    os.Getenv("PATH_TO_LOG_FILE"),
	}
}
