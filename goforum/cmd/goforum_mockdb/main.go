package main

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/webapi"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Could not load .env")
		return
	}

	jsonTokenSecret := os.Getenv("JSON_TOKEN_SECRET")

	if len(jsonTokenSecret) == 0 {
		log.Println("environment variable not found: JSON_TOKEN_SECRET")
		return;
	}

	a := core.App{
		ApiDriver: &webapi.WebApi{},
		Database:  mockdatabase.New(),
	}

	config := core.AppConfig{
		TokenSecret: jsonTokenSecret,
	}

	a.Init(config)
	a.Run()
}
