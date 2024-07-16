package main

import (
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi/v5"
)

func main() {
	routerChi := chi.NewRouter()
	routerChi.Get("/ping", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("pong"))
	})
	routerChi.Post("/json", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(map[string]string{"hello": "world"})
	})

	routerChi.Get("/greetings", greetings)
	http.ListenAndServe(":8080", routerChi)
}

type Greeting struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func greetings(responseWriter http.ResponseWriter, request *http.Request) {

	var greeting Greeting

	err := json.NewDecoder(request.Body).Decode(&greeting)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	responseWriter.Write([]byte("Hello, " + greeting.Name + " " + greeting.LastName))
}
