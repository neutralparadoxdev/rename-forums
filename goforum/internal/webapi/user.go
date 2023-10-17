package webapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountUser(router fiber.Router, app *core.App) {
	group := router.Group("/api/user")

	group.Post("/", func(c *fiber.Ctx) error {
		type NewUser struct {
			Username string `json:"username" form:"username"`
			Email    string `json:"email" form:"email"`
			Eula     bool   `json:"eula" form:"eula"`
			Password string `json:"password" form:"password"`
		}

		req := new(NewUser)

		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(400)
		}

		userManager := app.GetUserManager()

		user, err := userManager.CreateUser(req.Username, req.Email, req.Password, req.Eula)

		if err != nil {
			return c.SendStatus(400)
		}

		if user != nil {
			return c.SendStatus(201)
		} else {
			return c.SendStatus(400)
		}

	})
}
