package contexthelper

import (
	"context"
	"mercado-livre-integration/internal/infrastructure/log"
)

type contextKey int

const (
	loggerKey contextKey = iota
)

func GetLogger(ctx context.Context) (logs.Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(logs.Logger)
	return logger, ok
}

func SetLogger(ctx context.Context, logger logs.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
