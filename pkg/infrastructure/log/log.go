package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	With(args ...interface{})
	Shutdown()
}

type logger struct {
	logger *zap.SugaredLogger
}

func (l *logger) Shutdown() {
	l.logger.Sync()
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.logger.Infow(msg, args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.logger.Errorw(msg, args...)
}

func (l *logger) Warning(msg string, args ...interface{}) {
	l.logger.Warnw(msg, args...)
}

func (l *logger) With(args ...interface{}) {
	l.logger = l.logger.With(args...)
}

func New(appName string) Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "timestamp"
	l, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	l = l.WithOptions(zap.WithCaller(false))
	l = l.With(zap.String("tag", appName))
	return &logger{
		logger: l.Sugar(),
	}
}
