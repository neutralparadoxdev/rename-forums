package webapi

import (
	"github.com/gofiber/fiber/v2" // swagger handler
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type WebApi struct {
	fiberApp *fiber.App
	core     *core.App
}

func (w *WebApi) Init(app *core.App) error {
	w.fiberApp = fiber.New()

	w.core = app

	MountSession(w.fiberApp, app)
	MountUser(w.fiberApp, app)
	MountForum(w.fiberApp, app)

	return nil
}

func (w *WebApi) Run() error {
	w.fiberApp.Listen(":3001")
	return nil
}
