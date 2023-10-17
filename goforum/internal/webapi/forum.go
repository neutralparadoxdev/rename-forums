package webapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountForum(router fiber.Router, app *core.App) {
	group := router.Group("/api/forum")

	group.Post("/", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		jwtString, exists := headers["Bearer-Token"]

		var session core.Session

		if exists {
			session = core.NewSession(jwtString)
			ok, err := app.GetSessionManager().VerifySession(&session)

			if err != nil {
				return c.SendStatus(500)
			}

			if !ok {
				return c.SendStatus(400)
			}
		}

		type NewForumRequest struct {
			Title       string `json:"title" form:"title"`
			Description string `json:"description" form:"description"`
		}

		req := new(NewForumRequest)

		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(400)
		}

		ownerId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(400)
		}

		err = app.GetForumManager().CreateForum(req.Title, req.Description, ownerId)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(201)
	})

	group.Get("/", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		jwtString, exists := headers["Bearer-Token"]

		var session core.Session

		type ForumReponse struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		if exists {
			session = core.NewSession(jwtString)
			ok, err := app.GetSessionManager().VerifySession(&session)

			if err != nil {
				return c.SendStatus(500)
			}

			if !ok {
				return c.SendStatus(400)
			}

			userId, err := session.GetUserId()

			if err != nil {
				return c.SendStatus(500)
			}

			listOfForums, err := app.GetForumManager().GetAll(&userId)
			if err != nil {
				return c.SendStatus(500)
			}

			response := make([]ForumReponse, 0)
			for _, v := range listOfForums {
				response = append(response, ForumReponse{
					Title:       v.Title,
					Description: v.Description,
				})
			}

			return c.JSON(response)
		}

		listOfForums, err := app.GetForumManager().GetAll(nil)
		if err != nil {
			return c.SendStatus(500)
		}

		response := make([]ForumReponse, 0)
		for _, v := range listOfForums {
			response = append(response, ForumReponse{
				Title:       v.Title,
				Description: v.Description,
			})
		}

		return c.JSON(response)
	})
}
