package main

import (
	"api/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		return
	}

	app.Launch()
}
