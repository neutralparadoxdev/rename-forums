package webapi

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountMe(router fiber.Router, app *core.App) {
	group := router.Group("/api/me")

	group.Get("/", func(c *fiber.Ctx) error {
		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(401)
			} else {
				return c.SendStatus(500)
			}
		}

		if session == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		username, err := session.GetUsername()

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		log.Printf("username requested: %s", username)

		user, err := app.GetUserManager().GetUserByName(username)

		if err != nil {
			return c.SendStatus(500)
		}

		type UserResponse struct {
			CreatedAt    string `json:"created_at"`
			LastLogin    string `json:"last_login"`
			LastModified string `json:"last_modified"`
			UserName     string `json:"username"`
		}

		res := UserResponse{
			CreatedAt:    user.CreatedAt.Format(time.UnixDate),
			LastLogin:    user.LastLogin.Format(time.UnixDate),
			LastModified: user.LastModified.Format(time.UnixDate),
			UserName:     user.Username,
		}

		return c.JSON(res)
	})
}
