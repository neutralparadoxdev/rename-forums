package webapi

import (
	"fmt"
	"log"
	"time"

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

	group.Get("/:forumName", func(c *fiber.Ctx) error {
		forumName := c.Params("forumName")

		forum, err := app.GetForumManager().GetForum(forumName)

		if err != nil {
			return c.SendStatus(500)
		}

		if forum == nil {
			log.Printf("Forum Not found")
			return c.SendStatus(404)
		}

		posts, err := app.GetPostManager().GetPosts(forumName)

		if err != nil {
			return c.SendStatus(500)
		}

		type PostDTO struct {
			Title      string    `json:"title"`
			AuthorName string    `json:"authorName"`
			CreatedAt  time.Time `json:"createdAt"`
			Id         string    `json:"id"`
		}

		type CompleteForumDTO struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Posts       []PostDTO `json:"posts"`
		}

		postsdto := make([]PostDTO, 0)
		for _, v := range posts {
			postsdto = append(postsdto, PostDTO{
				Title:      v.Title,
				AuthorName: v.AuthorName,
				CreatedAt:  v.CreatedAt,
				Id:         fmt.Sprintf("%d", v.Id),
			})
		}

		forumResponse := CompleteForumDTO{
			Title:       forum.Title,
			Description: forum.Description,
			Posts:       postsdto,
		}

		return c.JSON(forumResponse)
	})

	group.Get("/", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		jwtString, exists := headers["Bearer-Token"]

		var session core.Session

		type ForumReponseDTO struct {
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

			response := make([]ForumReponseDTO, 0)
			for _, v := range listOfForums {
				response = append(response, ForumReponseDTO{
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

		response := make([]ForumReponseDTO, 0)
		for _, v := range listOfForums {
			response = append(response, ForumReponseDTO{
				Title:       v.Title,
				Description: v.Description,
			})
		}

		return c.JSON(response)
	})
}
