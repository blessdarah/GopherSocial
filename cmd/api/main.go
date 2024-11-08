package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &application{
		config: config{
			addr: os.Getenv("API_PORT"),
		},
	}
	mux := app.mount()
	log.Fatal(app.serve(mux))
}
