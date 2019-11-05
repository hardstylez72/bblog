package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New(env string) *logrus.Logger {
	log := logrus.New()

	if env != "dev" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}

	log.SetOutput(os.Stdout)
	return log
}
