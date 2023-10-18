package webapi

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountGroup(router fiber.Router, app *core.App) {
	group := router.Group("/api/group")

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

		var userId int64

		valid := false

		if session != nil {
			var err error

			userId, err = session.GetUserId()

			if err != nil {
				return c.SendStatus(500)
			}

			valid = true

		}

		postMan := app.GetPostManager()
		var posts []core.Post
		var err error
		if valid {
			posts, err = postMan.GetAllPosts(&userId)
		} else {
			posts, err = postMan.GetAllPosts(nil)
		}

		if err != nil {
			return c.SendStatus(500)
		}

		type PostDTO struct {
			Title      string    `json:"title"`
			Forum      string    `json:"forum"`
			AuthorName string    `json:"authorName"`
			CreatedAt  time.Time `json:"createdAt"`
			Id         string    `json:"id"`
		}

		postsdto := make([]PostDTO, 0)
		for _, v := range posts {
			postsdto = append(postsdto, PostDTO{
				Title:      v.Title,
				AuthorName: v.AuthorName,
				CreatedAt:  v.CreatedAt,
				Forum:      v.ForumPostedName,
				Id:         fmt.Sprintf("%d", v.Id),
			})
		}

		return c.JSON(postsdto)
	})

}
