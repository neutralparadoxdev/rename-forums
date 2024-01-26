package main

import (
	"log"
	"os"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/postgresdb"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/webapi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env")
		return
	}

	db, err := postgresdb.New(os.Getenv("DATABASE_URL"))
	defer db.Close()

	if err != nil {
		log.Printf("Could not create database %v", err)
	}

	jsonTokenSecret := os.Getenv("JSON_TOKEN_SECRET")

	if len(jsonTokenSecret) == 0 {
		log.Println("environment variable not found: JSON_TOKEN_SECRET")
		return;
	}

	a := core.App{
		ApiDriver: &webapi.WebApi{},
		Database:  db,
	}

	config := core.AppConfig{
		TokenSecret: jsonTokenSecret,
	}

	a.Init(config)
	a.Run()
}
