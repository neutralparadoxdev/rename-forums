package webapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountSession(router fiber.Router, app *core.App) {
	group := router.Group("/api/session")

	group.Post("/new", func(c *fiber.Ctx) error {
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

}
