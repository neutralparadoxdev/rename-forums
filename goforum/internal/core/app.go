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
	if app.ApiDriver != nil {
		if err := app.ApiDriver.Init(app); err != nil {
			log.Fatal(err)
			return err
		}
	}

	app.Authenticator = NewAuthenticator()
	app.UserManager = NewUserManager(app.Database, app.Authenticator)
	app.sessionManager = NewSessionManager(config.TokenSecret, app.UserManager, app.Authenticator, app.Database.GetSessionRepository())
	return nil
}

func (app *App) Run() {
	if app.ApiDriver != nil {
		app.ApiDriver.Run()
	}
}

func (app *App) GetSessionManager() *SessionManager {
	return app.sessionManager
}

func (app *App) GetUserManager() *UserManager {
	return app.UserManager
}
