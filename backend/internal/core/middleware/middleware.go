package middleware

import (
	"net/http"
	"strings"
	"time"

	"actualiteco/actualite/internal/core/logger"
	"actualiteco/actualite/internal/core/transport/response"

	"go.uber.org/zap"
)

var skipLogPrefixes = []string{"/swagger/", "/swagger"}

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) == 0 {
		return h
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func shouldSkipLogging(r *http.Request) bool {
	if r.Method == http.MethodOptions {
		return true
	}

	path := r.URL.Path

	for _, p := range skipLogPrefixes {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

func LogMiddleware(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if shouldSkipLogging(r) {
				next.ServeHTTP(w, r)
				return
			}

			log := log.With(zap.String("URL", r.URL.String()))

			ctx := logger.InContext(r.Context(), log)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func PanicMiddleware(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler := response.NewResponseHandler(log, w)

			defer func() {
				if err := recover(); err != nil {
					log.Error("panic happened", zap.Any("err", err))
					handler.PanicResponse(err, "panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func TraceMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if shouldSkipLogging(r) {
				next.ServeHTTP(w, r)
				return
			}

			log := logger.FromContext(r.Context())

			before := time.Now()
			log.Info(">>> incoming request", zap.String("METHOD", r.Method))

			next.ServeHTTP(w, r)

			after := time.Now()

			log.Info("<<< request handled", zap.Duration("latency", time.Duration(after.Sub(before))))
		})
	}
}

func CORSMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := map[string]struct{}{
				"http://localhost:3000": {},
				"http://localhost:8080": {},
			}

			origin := r.Header.Get("Origin")

			if _, ok := allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
