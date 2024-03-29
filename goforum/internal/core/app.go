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
	ForumManager   *ForumManager
	PostManager    *PostManager
	VoteManager    *VoteManager
	CommentManager *CommentManager
}

func (app *App) Init(config AppConfig) error {
	log.Print("App::Init: Initialization Starting")
	if app.ApiDriver != nil {
		if err := app.ApiDriver.Init(app); err != nil {
			log.Fatal(err)
			return err
		}
	}

	app.Authenticator = NewAuthenticator()
	app.UserManager = NewUserManager(app.Database, app.Authenticator)
	app.sessionManager = NewSessionManager(config.TokenSecret, app.UserManager, app.Authenticator, app.Database.GetSessionRepository())
	app.ForumManager = NewForumManager(app.Database)
	app.PostManager = NewPostManager(app.Database)
	app.VoteManager = NewVoteManager(app.Database)
	app.CommentManager = NewCommentManager(app.Database)

	log.Print("App::Init: Initialization Done")
	return nil
}

func (app *App) Run() {
	if app.ApiDriver != nil {
		log.Println("App::Run: ApiDriver Running")
		app.ApiDriver.Run()
	}
}

func (app *App) GetSessionManager() *SessionManager {
	return app.sessionManager
}

func (app *App) GetUserManager() *UserManager {
	return app.UserManager
}

func (app *App) GetForumManager() *ForumManager {
	return app.ForumManager
}

func (app *App) GetPostManager() *PostManager {
	return app.PostManager
}

func (app *App) GetVoteManager() *VoteManager {
	return app.VoteManager
}

func (app *App) GetCommentManager() *CommentManager {
	return app.CommentManager
}
