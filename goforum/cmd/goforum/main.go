package main

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/webapi"
)

func main() {
	a := core.App{
		ApiDriver: &webapi.WebApi{},
	}

	a.Init()
	a.Run()

}
