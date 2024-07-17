package middleware

import (
	"log"
	"net/http"
	"web/config"
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
		log.Println("Request received")
		next.ServeHTTP(w, r)
	})
}
