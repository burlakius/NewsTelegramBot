package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	TGbotToken          string
	LogLevel            logrus.Level
	LogPath             string
	RedisAddress        string
	RedisPort           string
	MariaDBRootPassword string
	MariaDBDatabase     string
	MariaDBUser         string
	MariaDBPassword     string
)

func LoadConfig() {
	err := godotenv.Load()
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

	TGbotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	LogLevel = logLevels[level]
	LogPath = os.Getenv("PATH_TO_LOG_FILE")
	RedisAddress = os.Getenv("REDIS_ADDRESS")
	RedisPort = os.Getenv("REDIS_PORT")
	MariaDBRootPassword = os.Getenv("MARIADB_ROOT_PASSWORD")
	MariaDBDatabase = os.Getenv("MARIADB_DATABASE")
	MariaDBUser = os.Getenv("MARIADB_USER")
	MariaDBPassword = os.Getenv("MARIADB_PASSWORD")
}
