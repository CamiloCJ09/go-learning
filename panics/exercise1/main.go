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

	_, err := os.Open("filename.ext")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
}
