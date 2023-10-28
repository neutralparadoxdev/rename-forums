package webapi

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountPost(router fiber.Router, app *core.App) {
	group := router.Group("/api/post/")

	group.Post("/:forumName/:id/upvote", func(c *fiber.Ctx) error {
		stringId := c.Params("id")

		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			return c.SendStatus(404)
		}

		//forumName := c.Params("forumName")

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				return c.SendStatus(500)
			}
		}

		userId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		app.GetVoteManager().ChangeVote(userId, id, 1)

		return c.SendStatus(202)
	})

	group.Post("/:forumName/:id/downvote", func(c *fiber.Ctx) error {
		stringId := c.Params("id")

		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			return c.SendStatus(404)
		}

		//forumName := c.Params("forumName")

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				return c.SendStatus(500)
			}
		}

		userId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		app.GetVoteManager().ChangeVote(userId, id, -1)

		return c.SendStatus(202)
	})

	group.Delete("/:forumName/:id/deletevote", func(c *fiber.Ctx) error {
		stringId := c.Params("id")

		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			return c.SendStatus(404)
		}

		//forumName := c.Params("forumName")

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				return c.SendStatus(500)
			}
		}

		userId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		app.GetVoteManager().ChangeVote(userId, id, 0)

		return c.SendStatus(202)
	})

	group.Get("/:forumName/:id", func(c *fiber.Ctx) error {
		stringId := c.Params("id")

		id, err := strconv.ParseInt(stringId, 10, 64)
		if err != nil {
			return c.SendStatus(404)
		}

		forumName := c.Params("forumName")

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(fiber.StatusUnauthorized)
			} else {
				return c.SendStatus(500)
			}
		}

		if session != nil {
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
				if err.Error() == "user_cant_post" {
					return c.SendStatus(401)
				}
				return c.SendStatus(500)
			}

			if post == nil {
				return c.SendStatus(404)
			}

			type PostDTO struct {
				Title      string    `json:"title"`
				Body       string    `json:"body"`
				AuthorName string    `json:"authorName"`
				CreatedAt  time.Time `json:"createdAt"`
			}

			postdto := PostDTO{
				Title:      post.Title,
				Body:       post.Body,
				AuthorName: post.AuthorName,
				CreatedAt:  post.CreatedAt,
			}

			return c.JSON(postdto)
		}
	})

	group.Post("/", func(c *fiber.Ctx) error {

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Printf("post_post: %s", webErr)
			if webErr.Code != WebApiErrorServerError.Code {
				return c.SendStatus(401)
			} else {
				return c.SendStatus(500)
			}
		}

		if session == nil {
			return c.SendStatus(401)
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

		if len(req.Title) == 0 || len(req.Body) == 0 || len(req.ForumName) == 0 {
			return c.SendStatus(400)
		}

		userId, err := session.GetUserId()

		if err != nil {
			return c.SendStatus(500)
		}

		postId, err := app.GetPostManager().CreatePost(req.Title, req.Body, req.ForumName, userId)

		if err != nil {
			return c.SendStatus(500)
		}
		type PostIdReponse struct {
			Id string `json:"id" form:"id"`
		}

		response := PostIdReponse{
			Id: fmt.Sprintf("%d", postId),
		}

		return c.JSON(response)
	})
}
