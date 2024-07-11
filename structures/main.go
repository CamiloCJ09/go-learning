package main

import (
	"encoding/json"
	"fmt"

	"github.com/CamiloCJ09/go-learning/structures/model"
)

func main() {
	p1 := model.Product{Id: 1, Name: "Product 1", Price: 12.99, Description: "Description 1", Category: "Category 1"}
	p2 := model.Product{Id: 2, Name: "Product 2", Price: 24.99, Description: "Description 2", Category: "Category 2"}
	p3 := model.Product{Id: 3, Name: "Product 3", Price: 36.99, Description: "Description 3", Category: "Category 3"}

	p1.Save()
	p2.Save()
	p3.Save()

	product := model.GetById(2)
	productJson, _ := json.MarshalIndent(product, "", "  ")
	fmt.Println(string(productJson))

	allProducts, _ := json.MarshalIndent(p1.GetAll(), "", "  ")
	fmt.Println(string(allProducts))

	e1 := model.Employee{Id: 1, Position: "Position 1", Person: model.Person{Id: 1, Name: "Person 1", DateOfBirth: "01/01/2000"}}
	e1.PrintEmployee()
}

// Product is a struct that represents a product
