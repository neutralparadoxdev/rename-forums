package core

type ApiDriver interface {
	Init(app *App) error
	Run() error
}
