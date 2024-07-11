package main

import (
	"errors"
	"fmt"
)

var clients []*Client

func main() {
	NewClient("clients.txt", "John Doe", "123456789", "123 Main St", 1)
	NewClient("clients.txt", "Jane Doe", "987654321", "321 Main St", 2)
	NewClient("clients.txt", "John Smith", "123456789", "123 Main St", 1)
	fmt.Println("Execution finished")
}

type Client struct {
	File        string
	Name        string
	Id          int
	PhoneNumber string
	Address     string
}

func NewClient(file string, name string, phoneNumber string, address string, id int) *Client {

	clientExist, existingError := validateExistingClient(id)

	if clientExist {

		handlePanic("Client already exists")
	}

	invalidParameters, parametersError := validateClientFields(file, name, phoneNumber, address, id)

	if invalidParameters {

		handlePanic("Invalid client fields")
	}

	if existingError == nil && parametersError == nil {
		client := &Client{
			File:        file,
			Name:        name,
			Id:          id,
			PhoneNumber: phoneNumber,
			Address:     address,
		}

		clients = append(clients, client)
		fmt.Println("Client created successfully")
		return client
	} else {
		handlePanic("Several errors were detected at runtime")
	}
	return nil
}

func validateExistingClient(id int) (bool, error) {
	for _, client := range clients {
		if client.Id == id {
			return true, errors.New("Client exist")
		}
	}
	return false, nil
}

func validateClientFields(file, name, phoneNumber, address string, id int) (bool, error) {
	if file == "" || name == "" || phoneNumber == "" || address == "" || id == 0 {
		return true, errors.New("invalid client fields")
	}
	return false, nil
}

func handlePanic(msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(msg)
}
