package main

import (
	"log"
	"os"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/postgresdb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env")
		return
	}

	db, err := postgresdb.New(os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Printf("Could not create database %v", err)
	}

	defer db.Close()
}
