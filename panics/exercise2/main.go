package main

import (
	"fmt"
	"os"
)

func main() {

	openFile()
	fmt.Println("Execution finished")
}

func openFile() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Open("costumers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	b1 := make([]byte, 5)
	_, errRead := file.Read(b1)
	if errRead != nil {
		panic("Error reading file")
	}
	fmt.Println(string(b1))
	file.Close()
}
