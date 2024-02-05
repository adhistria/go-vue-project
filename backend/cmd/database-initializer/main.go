package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/adhistria/internal/database"

	// "github.com/golang-migrate/migrate/database"
	"github.com/adhistria/backend/go-vue-project/internal/database"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("dbHost,dbDriver,dbUser,dbPassword,dbName,dbPort", dbHost, dbDriver, dbUser, dbPassword, dbName, dbPort)

	// connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)

	db, err := sqlx.Connect(dbDriver, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	database.InitializeDB(db)

}
