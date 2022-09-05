package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// library for conenct postgresql
	_ "github.com/lib/pq"
)

var (
	// DB variable for connection DB postgresql
	DB *sql.DB
)

func init() {
	fmt.Println("Start connecting DB")
	err := Connect()
	if err != nil {
		fmt.Print("Error connecting DB", err)
	}

	err = CreateTables()
	if err != nil {
		fmt.Println(err)
	}
}

// Connect function for checking connection to postgresql
func Connect() error {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error connecting: ", err)
		return err
	}

	err = result.Ping()
	if err != nil {
		log.Println("Error DB Ping : ", err)
		return err
	}

	DB = result
	return nil
}
