package main

import (
	"fmt"

	"github.com/CamiloCJ09/go-learning/finalProject/model"
	"github.com/CamiloCJ09/go-learning/finalProject/utils"
)

func main() {
	readAndStoreData()

}

func readAndStoreData() {
	data, err := utils.FromCSVtoTickets()
	fmt.Println(data, err)
}

func mockedData() {
	ticket := model.Ticket{
		Id:          1,
		Name:        "John Doe",
		Email:       "",
		Destination: "New York",
		Country:     "USA",
		FlightTime:  "2021-12-01T08:00:00Z",
		Price:       1000,
	}
	ticket.CreateNewTicket()
	ticket = model.Ticket{
		Id:          2,
		Name:        "Jane Doe",
		Email:       "",
		Destination: "Paris",
		Country:     "France",
		FlightTime:  "2021-12-01T08:00:00Z",
		Price:       1200,
	}

}
