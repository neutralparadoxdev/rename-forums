package webapi

import (
	"github.com/gofiber/fiber/v2"
)

type WebApi struct {
	fiberApp *fiber.App
}

func (w *WebApi) Init() error {
	w.fiberApp = fiber.New()

	w.fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	return nil
}

func (w *WebApi) Run() error {
	w.fiberApp.Listen(":3001")
	return nil
}
