package main

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	mockdatabase "github.com/neutralparadoxdev/rename-forums/goforum/internal/mock_database"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/webapi"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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
