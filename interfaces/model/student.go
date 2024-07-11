package model

import "fmt"

type Student struct {
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	DNI          string `json:"dni"`
	StartingDate string `json:"starting_date"`
}

var students []Student

func (s Student) Save() {
	students = append(students, s)
}

func (s Student) PrintInformation() {
	for _, student := range students {
		fmt.Println("Name:", student.Name)
		fmt.Println("Last Name:", student.LastName)
		fmt.Println("DNI:", student.DNI)
		fmt.Println("Starting Date:", student.StartingDate)
		fmt.Println("-------------------------------")
	}
}
