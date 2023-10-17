package webapi

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountPost(router fiber.Router, app *core.App) {
	group := router.Group("/api/post/")

	group.Get("/:forumName/:id", func(c *fiber.Ctx) error {
		stringId := c.Params("id")

		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			return c.SendStatus(404)
		}

		forumName := c.Params("forumName")

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

			userId, err := session.GetUserId()

			if err != nil {
				return c.SendStatus(500)
			}

			post, err := app.GetPostManager().GetPost(id, forumName, &userId)

			if err != nil {
				return c.SendStatus(500)
			}

			if post == nil {
				return c.SendStatus(404)
			}

			return c.JSON(post)
		} else {
			post, err := app.GetPostManager().GetPost(id, forumName, nil)

			if err != nil {
				return c.SendStatus(500)
			}

			if post == nil {
				return c.SendStatus(404)
			}

			return c.JSON(post)
		}
	})

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

		type NewPostRequest struct {
			Title     string `json:"title" form:"title"`
			Body      string `json:"body" form:"body"`
			ForumName string `json:"forumName" form:"forumName"`
		}

		req := new(NewPostRequest)

		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(400)
		}

		userId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(500)
		}

		ok, err := app.GetPostManager().CreatePost(req.Title, req.Body, req.ForumName, userId)

		if err != nil {
			return c.SendStatus(500)
		}

		if ok {
			return c.SendStatus(201)
		} else {
			return c.SendStatus(500)
		}
	})
}
