package initialize

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
