package main

import (
	"fmt"
	"time"
)

func main() {

	salsa := make(chan string)
	pasta := make(chan string)

	go MakePasta(pasta)
	go MakeSalsa(salsa)

	pastaDone := <-pasta
	salsaDone := <-salsa

	fmt.Println(pastaDone)
	fmt.Println(salsaDone)
}

func MakePasta(chanel chan<- string) {
	startTime := time.Now()

	fmt.Println("Boiling water")
	time.Sleep(1 * time.Second)
	fmt.Println("Add pasta")
	time.Sleep(1 * time.Second)
	fmt.Println("Cooking pasta")
	time.Sleep(1 * time.Second)
	fmt.Println("Pasta is ready!")

	fmt.Println("Time taken to make pasta: ", time.Since(startTime))

	chanel <- "Pasta is ready!"
}

func MakeSalsa(chanel chan<- string) {
	startTime := time.Now()

	fmt.Println("Chopping tomatoes")
	time.Sleep(1 * time.Second)
	fmt.Println("Chopping onions")
	time.Sleep(1 * time.Second)
	fmt.Println("Chopping chillies")
	time.Sleep(1 * time.Second)
	fmt.Println("Mixing everything")
	time.Sleep(1 * time.Second)
	fmt.Println("Salsa is ready!")

	fmt.Println("Time taken to make salsa: ", time.Since(startTime))

	chanel <- "Salsa is ready!"
}
