package main

import (
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

	dbDriver := os.Getenv("DB_DRIVER")
	db, err := sqlx.Connect(dbDriver, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	database.InitializeDB(db)

}
