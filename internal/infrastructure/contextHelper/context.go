package contexthelper

import (
	"context"
	"mercado-livre-integration/internal/infrastructure/log"
)

type contextKey int

const (
	loggerKey contextKey = iota
	requestIDKey
)

func GetLogger(ctx context.Context) (logs.Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(logs.Logger)
	return logger, ok
}

func SetLogger(ctx context.Context, logger logs.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return ""
	}
	return requestID
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
