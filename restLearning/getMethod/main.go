package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CamiloCJ09/go-learning/restLearning/getMethod/model"
	"github.com/CamiloCJ09/go-learning/restLearning/getMethod/utils"
	"github.com/go-chi/chi"
)

func main() {
	utils.ReadFromJsonFile()
	StartServer()
}

func StartServer() *chi.Mux {
	routerChi := chi.NewRouter()

	routerChi.Route("/api/v1", func(outerRouter chi.Router) {

		outerRouter.Route("/products", func(innerRouter chi.Router) {

			innerRouter.Get("/", GetProducts)
			innerRouter.Get("/{id}", GetProductById)
			innerRouter.Get("/search", SearchProduct)
			innerRouter.Post("/", CreateProduct)
		})

		outerRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to the server"))
		})

	})

	http.ListenAndServe(":8080", routerChi)
	return routerChi
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := model.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
	// Get all products
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product := model.GetProductById(intId)
	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	price := r.URL.Query().Get("priceGt")
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products := model.SearchProductByPrice(floatPrice)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	productDTO := model.RequestProduct{}
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product := model.CreateProduct(productDTO.Name, productDTO.Quantity, productDTO.Code_value, productDTO.IsPublished, productDTO.Expiration, productDTO.Price)
	if product.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}
