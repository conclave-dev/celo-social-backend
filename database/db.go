package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// Required by sqlx as empty import
	_ "github.com/lib/pq"
)

// DBClient holds the instance of the connected database
var DBClient *sqlx.DB

// SetupDatabaseConnection initiates the connection to the Postgresql database
func SetupDatabaseConnection() {
	// Construct the db connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSL"),
	)

	client, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	DBClient = client
}
