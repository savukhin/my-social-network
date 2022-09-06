package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("FROM init", os.Getenv("POSTGRES_HOST"))
}
