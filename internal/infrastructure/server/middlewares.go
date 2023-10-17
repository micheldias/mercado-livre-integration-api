package server

import (
	contexthelper "mercado-livre-integration/internal/infrastructure/contextHelper"
	logs "mercado-livre-integration/internal/infrastructure/log"
	"net/http"
)

func InjectLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logs.New("mercado-livre-api")
		ctx = contexthelper.SetLogger(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))

		//http.Error(w, "Forbidden", http.StatusForbidden)

	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logs.New("mercado-livre-api")
		ctx = contexthelper.SetLogger(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))

		//http.Error(w, "Forbidden", http.StatusForbidden)

	})
}
