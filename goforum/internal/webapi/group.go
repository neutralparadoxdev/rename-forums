package webapi

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountGroup(router fiber.Router, app *core.App) {
	group := router.Group("/api/group")

	group.Get("/", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		jwtString, exists := headers["Bearer-Token"]

		var session core.Session

		var userId int64

		valid := false

		if exists {
			session = core.NewSession(jwtString)
			ok, err := app.GetSessionManager().VerifySession(&session)

			if err != nil {
				return c.SendStatus(500)
			}

			if !ok {
				return c.SendStatus(400)
			}

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
