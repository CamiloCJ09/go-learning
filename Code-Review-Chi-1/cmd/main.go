package main

import (
	"app/cmd/server"
	"fmt"
	"log"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "../docs/db/vehicles_100.json",
	}
	app := server.NewServerChi(cfg)
	// - run
	log.Println("Server is running on", cfg.ServerAddress)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
