package model

import (
	"fmt"
	"time"
)

var Products []Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code_value  string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type RequestProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code_value  string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func GetProducts() []Product {
	return Products
}

func GetProductById(id int) Product {
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}

func SearchProductByPrice(price float64) []Product {
	products := GetProducts()
	result := []Product{}
	for _, product := range products {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}

func CreateProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64) Product {
	id := len(Products) + 1

	if verifyProductIdExists(id) {
		fmt.Println("Product ID already exists")
		return Product{}
	}

	if !verifyNotEmptyParameters(name, quantity, code_value, expiration, price) {
		fmt.Println("Empty parameters")
		return Product{}
	}

	if verifyCodeValueExists(code_value) {
		fmt.Println("Code value already exists")
		return Product{}
	}

	if !verifyExpirationDate(expiration) {
		fmt.Println("Invalid expiration date")
		return Product{}
	}

	productToAdd := Product{
		ID:          id,
		Name:        name,
		Quantity:    quantity,
		Code_value:  code_value,
		IsPublished: is_published,
		Expiration:  expiration,
		Price:       price,
	}
	Products = append(Products, productToAdd)
	return productToAdd
}

func verifyProductIdExists(id int) bool {
	for _, product := range Products {
		if product.ID == id {
			return true
		}
	}
	return false
}

func verifyNotEmptyParameters(name string, quantity int, code_value string, expiration string, price float64) bool {
	if name == "" || quantity == 0 || code_value == "" || expiration == "" || price == 0 {
		return false
	}
	return true
}

func verifyCodeValueExists(code_value string) bool {
	for _, product := range Products {
		if product.Code_value == code_value {
			return true
		}
	}
	return false
}

func verifyExpirationDate(expiration string) bool {
	timeExpiration, err := time.Parse("02/01/2006", expiration)
	if err != nil {
		return false
	}
	if timeExpiration.Before(time.Now()) {
		return false
	}

	return true
}
