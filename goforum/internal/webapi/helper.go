package webapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func containsI64(container *[]int64, item int64) bool {
	for _, val := range *container {
		if val == item {
			return true
		}
	}
	return false
}

func OptionalSessionAuth(
	app *core.App,
	handle func(*fiber.Ctx, *core.Session) error) func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		session, err := CheckForSession(c, app.GetSessionManager())

		if err != nil {
			log.Printf("Error Code: %d", err.Code)
			if err.Code == WebApiErrorTokenValidationFailed.Code {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				if err.Code == WebApiErrorTokenInvalid.Code {
					return c.SendStatus(fiber.StatusUnauthorized)
				}
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		return handle(c, session)
	}
}

func RequiredAuth(
	app *core.App,
	handle func(*fiber.Ctx, *core.Session) error) func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		session, err := CheckForSession(c, app.GetSessionManager())

		if err != nil {
			log.Printf("Error Code: %d", err.Code)
			if err.Code == WebApiErrorTokenValidationFailed.Code {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				if err.Code == WebApiErrorTokenInvalid.Code {
					return c.SendStatus(fiber.StatusUnauthorized)
				}
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		if session == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return handle(c, session)
	}

}
