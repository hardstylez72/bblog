package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func New() *logrus.Logger {
	log := logrus.New()

	if viper.GetString("env") != "dev" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{})
	}

	log.SetOutput(os.Stdout)
	return log
}
