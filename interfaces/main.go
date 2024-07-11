package main

import (
	"fmt"

	"github.com/CamiloCJ09/go-learning/interfaces/model"
)

func main() {
	s1 := model.Student{Name: "Camilo", LastName: "CJ", DNI: "123456789", StartingDate: "01/01/2021"}
	s2 := model.Student{Name: "Juan", LastName: "Perez", DNI: "987654321", StartingDate: "01/01/2021"}

	s1.Save()
	s2.Save() // >

	s1.PrintInformation()

	p1 := model.CreateProduct(100.0, "Small")
	p2 := model.CreateProduct(200.0, "Medium")
	p3 := model.CreateProduct(300.0, "Large")

	fmt.Println(p1.GetPrice())
	fmt.Println(p2.GetPrice())
	fmt.Println(p3.GetPrice())
}
