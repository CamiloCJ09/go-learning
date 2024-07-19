package middleware

import (
	"app/cmd/config"
	"log"
	"net/http"
)

const (
	TOKEN = "TOKEN"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := config.NewConfig()
		token := config.Get(TOKEN)

		if token == "" {
			log.Println("Token not found in .env file")
			http.Error(w, "Token not found", http.StatusUnauthorized)
			return
		}

		if r.Header.Get("Authorization") != token {
			log.Println("Unauthorized access")
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Request received ", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Println("Request completed", r.Method, r.URL.Path, r.RemoteAddr)
	})
}
