package auth

import (
	"net/http"
)

func New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("user") == "otto" {
			next.ServeHTTP(wr, r)
		} else {
			wr.Write([]byte("500 error"))
		}
	})
}

