package main

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/webapi"
)

func main() {
	a := core.App{
		ApiDriver: &webapi.WebApi{},
		Database:  &mockdatabase.MockDatabase{},
	}

	config := core.AppConfig{
		TokenSecret: "MySuperSecretToken",
	}

	a.Init(config)
	a.Run()
}
