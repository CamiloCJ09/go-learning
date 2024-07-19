package main

import (
	"log"
	"net/http"
	"web/cmd/server/handler"
	"web/cmd/server/middleware"
	"web/config"
	"web/internal/product"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := startServer()

	config.LoadConfig()

	log.Println("Starting server on :2000...")
	err := http.ListenAndServe(":2000", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func startServer() *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ready for requests!"))
	})

	prodRepo, err := product.NewProductRepository("products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}

	prodService := product.NewProductService(prodRepo)

	prodHandler := handler.NewProductHandler(prodService)

	chiRouter.Route("/api/v1", func(apiRouter chi.Router) {

		apiRouter.Route("/products", func(prodRouter chi.Router) {

			prodRouter.Use(middleware.AuthMiddleware)

			prodRouter.Use(middleware.LogMiddleware)

			prodRouter.Get("/", prodHandler.GetAllProducts)

			prodRouter.Get("/{id}", prodHandler.GetProductByID)

			prodRouter.Get("/search", prodHandler.SearchProduct)

			prodRouter.Get("/consumer-price", prodHandler.GetConsumerPrice)

			prodRouter.Post("/", prodHandler.CreateProduct)

			prodRouter.Put("/{id}", prodHandler.UpdateProduct)

			prodRouter.Patch("/{id}", prodHandler.UpdateProduct)

			prodRouter.Delete("/{id}", prodHandler.DeleteProduct)

		})

	})

	return chiRouter
}
