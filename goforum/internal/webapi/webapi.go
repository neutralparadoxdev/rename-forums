package webapi

import (
	"github.com/gofiber/fiber/v2" // swagger handler
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"log"
)

type WebApi struct {
	fiberApp *fiber.App
	core     *core.App
}

func (w *WebApi) Init(app *core.App) error {
	log.Println("WebApi::Init: Initializing")
	w.fiberApp = fiber.New()

	w.core = app

	// Initialize default config (Assign the middleware to /metrics)
	w.fiberApp.Get("/metrics", monitor.New())

	// Or extend your config for customization
	// Assign the middleware to /metrics
	// and change the Title to `MyService Metrics Page`
	w.fiberApp.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	MountSession(w.fiberApp, app)
	MountUser(w.fiberApp, app)
	MountForum(w.fiberApp, app)
	MountPost(w.fiberApp, app)
	MountGroup(w.fiberApp, app)
	MountMe(w.fiberApp, app)
	MountComment(w.fiberApp, app)

	log.Println("WebApi::Init: Initialization Done")
	return nil
}

func (w *WebApi) Run() error {
	log.Println("WebApi::Run: Starting Listening")
	log.Println(w.fiberApp.Listen(":3001"))
	return nil
}
