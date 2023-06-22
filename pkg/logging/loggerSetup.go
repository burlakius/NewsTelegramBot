package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func LoggerSetup(pathToLogFile string, loggingLevel logrus.Level) {
	logFile, err := os.OpenFile(pathToLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, logFile)
		logrus.SetOutput(mw)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	logrus.SetLevel(loggingLevel)
}
