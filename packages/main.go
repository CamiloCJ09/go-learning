package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	PrintEnvVar()
	printResponse()
	copyContentFromFileToEmptyFile()
}

func PrintEnvVar() {
	fmt.Println(os.Getenv("MY_ENV_VAR"))
}

func printResponse() {
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/todos", nil)
	if err != nil {
		panic("The request could not be created")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("The request could not be executed")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Failed to read response body")
	}
	fmt.Println(string(body))
}

// 3.Copiar el contenido de un archivo a otro vacio

func copyContentFromFileToEmptyFile() {
	originalFile, err := os.Open("original.txt")
	if err != nil {
		panic("The file could not be opened")
	}
	defer originalFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic("The file could not be created")
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		panic("The file could not be copied")
	}
}
