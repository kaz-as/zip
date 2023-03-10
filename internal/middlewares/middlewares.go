package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/kaz-as/zip/pkg/logger"
)

type Middleware func(http.Handler) http.Handler

func Chain(mws []Middleware) Middleware {
	mwsCopy := make([]Middleware, len(mws))
	copy(mwsCopy, mws)

	return func(h http.Handler) (res http.Handler) {
		res = h
		for i := len(mwsCopy) - 1; i >= 0; i-- {
			res = mwsCopy[i](res)
		}
		return
	}
}

//// Global

func Logger(l logger.Interface) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r != nil {
				l.Info("new request: %v \"%v\"", r.Method, r.RequestURI)
			} else {
				l.Warn("logger: request is nil")
			}
			next.ServeHTTP(w, r)
		})
	}
}

func Recoverer(l logger.Interface) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					l.Error("panic: %v", rec)
					jsonBody, _ := json.Marshal(map[string]string{
						"error": "There was an internal server error",
					})

					if w.Header() != nil {
						w.Header().Set("Content-Type", "application/json")
					}
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(jsonBody)
					if err != nil {
						l.Error("cannot write internal server error: %s", err)
					}
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

//// Local
