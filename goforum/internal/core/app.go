package core

import (
	"log"
)

type AppConfig struct {
	TokenSecret string
}

type App struct {
	ApiDriver      ApiDriver
	sessionManager *SessionManager
	UserManager    *UserManager
	Authenticator  *Authenticator
	Database       Database
}

func (app *App) Init(config AppConfig) error {
	if err := app.ApiDriver.Init(app); err != nil {
		log.Fatal(err)
		return err
	}

	app.Authenticator = &Authenticator{}
	app.UserManager = NewUserManager(app.Database)
	app.sessionManager = NewSessionManager(config.TokenSecret, app.UserManager, app.Authenticator, app.Database.GetSessionRepository())
	return nil
}

func (app *App) Run() {

	app.ApiDriver.Run()
}

func (app *App) GetSessionManager() *SessionManager {
	return app.sessionManager
}
