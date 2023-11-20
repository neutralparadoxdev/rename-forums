package webapi

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

func MountComment(router fiber.Router, app *core.App) {
	group := router.Group("/api/comment")

	group.Get("/:forum/:postId/:commentId", func(c *fiber.Ctx) error {
		forumId := c.Params("forum")

		postId, err := strconv.ParseInt(c.Params("postId"), 10, 64)

		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		commentId, err := strconv.ParseInt(c.Params("commentId"), 10, 64)

		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		session, webErr := CheckForSession(c, app.GetSessionManager())

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(401)
			} else {
				return c.SendStatus(500)
			}
		}

		comments, err := app.GetCommentManager().GetCommentWithUserSession(
			commentId,
			postId,
			forumId,
			session,
		)

		if err != nil {
			return c.SendStatus(404)
		}

		type CommentDTO struct {
			Text         string `json:"text"`
			CommentOwner string `json:"commentOwner"`
			PostOwner    string `json:"postOwner"`
			Id           string `json:"id"`
		}

		commentsDto := make([]CommentDTO, 0)

		for i := range comments {
			postOwner := ""
			commentOwner := ""

			if comments[i].PostOwner != nil {
				postOwner = fmt.Sprintf("%d", *comments[i].PostOwner)
			}

			if comments[i].CommentOwner != nil {
				commentOwner = fmt.Sprintf("%d", *comments[i].CommentOwner)
			}

			commentDto := CommentDTO{
				Text:         comments[i].Text,
				CommentOwner: commentOwner,
				PostOwner:    postOwner,
				Id:           fmt.Sprintf("%d", comments[i].Id),
			}
			commentsDto = append(commentsDto, commentDto)
		}

		return c.JSON(commentsDto)
	})
}
