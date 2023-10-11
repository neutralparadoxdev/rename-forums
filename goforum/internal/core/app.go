package core

import (
	"log"
)

type App struct {
	ApiDriver ApiDriver
}

func (app *App) Init() error {
	if err := app.ApiDriver.Init(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (app *App) Run() {

	app.ApiDriver.Run()
}
