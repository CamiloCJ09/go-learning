package utils

import (
	"encoding/json"
	"os"

	"github.com/CamiloCJ09/go-learning/restLearning/getMethod/model"
)

func ReadFromJsonFile() {
	file, err := os.Open("products.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	products := []model.Product{}
	err = decoder.Decode(&products)
	if err != nil {
		panic(err)
	}

	model.Products = products
}
