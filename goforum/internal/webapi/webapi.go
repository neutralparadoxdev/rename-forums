package webapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type WebApi struct {
	fiberApp *fiber.App
	core     *core.App
}

func (w *WebApi) Init(app *core.App) error {
	w.fiberApp = fiber.New()

	w.core = app

	w.fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	w.fiberApp.Post("/api/session/new", func(c *fiber.Ctx) error {
		type SessionRequest struct {
			Username string `json:"username" form:"username"`
			Password string `json:"password" form:"password"`
		}

		req := new(SessionRequest)

		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(400)
		}

		sessionManager := app.GetSessionManager()
		session, err := sessionManager.CreateSession(req.Username, req.Password)
		if err != nil {
			log.Print("Attempted unauthorized Access")
			return c.SendStatus(401)
		}

		return c.SendString(session.ToString())
	})

	return nil
}

func (w *WebApi) Run() error {
	w.fiberApp.Listen(":3001")
	return nil
}
