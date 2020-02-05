package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Load the environment variables
	err := godotenv.Load(".env")

	// Initiate the migrations procedure
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSL"),
	)

	migration, err := migrate.New("file://database/migrations", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err)
	}
}
