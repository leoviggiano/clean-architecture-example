package log

import (
	"github.com/sirupsen/logrus"
)

type Option func(logger *logrus.Logger)

func WithTextFormatter() Option {
	return func(logger *logrus.Logger) {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:            true,
			DisableLevelTruncation: true,
			DisableTimestamp:       true,
		})
	}
}

func NewLogger(level logrus.Level, options ...Option) logrus.FieldLogger {
	logger := logrus.New()
	logger.Level = level

	for _, option := range options {
		option(logger)
	}

	logger.Info("[Log] Started with success")
	return logger
}

func IfError(logger logrus.FieldLogger, err error) {
	if err != nil {
		logger.Error(err)
	}
}
