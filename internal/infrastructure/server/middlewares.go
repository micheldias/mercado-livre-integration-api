package server

import (
	"github.com/google/uuid"
	contexthelper "mercado-livre-integration/internal/infrastructure/contextHelper"
	logs "mercado-livre-integration/internal/infrastructure/log"
	"net/http"
)

func InjectLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logs.New("mercado-livre-api")
		if requestID := contexthelper.GetRequestID(ctx); requestID != "" {
			logger.With("requestID", requestID)
		}

		ctx = contexthelper.SetLogger(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func InjectRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		if requestIDHeader := r.Header.Get("X-Request-Id"); requestIDHeader != "" {
			requestID = requestIDHeader
		}

		ctx := r.Context()
		ctx = contexthelper.SetRequestID(ctx, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger, _ := contexthelper.GetLogger(ctx)
		defer func(log logs.Logger) {
			if err := recover(); err != nil {
				log.Error("panic occurred:", "error", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"internal server error"}`))
			}
		}(logger)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
