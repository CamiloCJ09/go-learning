package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/CamiloCJ09/go-learning/finalProject/model"
)

func ReadFromFile() ([]byte, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)
	return buffer, nil
}

func FromCSVtoTickets() ([]model.Ticket, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var tickets []model.Ticket
	for _, record := range records {
		id, intErr := strconv.ParseInt(record[0], 10, 32)
		price, err := strconv.ParseFloat(record[6], 64)
		if err != nil || intErr != nil {
			return nil, err
		}
		ticket := model.Ticket{
			Id:          int(id),
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			Country:     record[4],
			FlightTime:  record[5],
			Price:       float64(price),
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func WriteToFile(data []byte) error {
	file, err := os.Create("tickets.json")
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
