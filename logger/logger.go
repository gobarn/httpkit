package logger

import (
	"log"
	"net/http"
)

func New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		next.ServeHTTP(wr, r)
	})
}

