//go:generate go run github.com/golang/mock/mockgen -package=mocks -source=$GOFILE -destination=../../testdata/mocks/log.go
package log

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(log string)
	Infof(format string, args ...interface{})
	Error(err error)
	Errorf(format string, args ...interface{})
	Fatal(err error)
}

type logger struct {
	Logrus logrus.FieldLogger
}

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

func NewLogger(level logrus.Level, options ...Option) Logger {
	logrus := logrus.New()
	logrus.Level = level

	for _, option := range options {
		option(logrus)
	}

	logrus.Info("[Log] Started with success")
	return &logger{logrus}
}

func (l logger) Info(log string) {
	l.Logrus.Info(log)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Logrus.Infof(format, args)
}

func (l logger) Error(err error) {
	l.Logrus.Error(err)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Logrus.Errorf(format, args)
}

func (l logger) Fatal(err error) {
	l.Logrus.Fatal(err)
}
