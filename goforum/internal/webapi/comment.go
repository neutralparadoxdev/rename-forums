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

	group.Post("/:forum/:postId/:commentId", func(c *fiber.Ctx) error {
		//postId, err := strconv.ParseInt(c.Params("postId"), 10, 64)

		//if err != nil {
		//	return c.SendStatus(fiber.StatusNotFound)
		//}

		commentId, err := strconv.ParseInt(c.Params("commentId"), 10, 64)

		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		session, webErr := CheckForSession(c, app.GetSessionManager())

		type NewCommentRequest struct {
			Text string `json:"text" form:"text"`
		}

		req := new(NewCommentRequest)

		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(400)
		}

		if webErr != nil {
			log.Print(webErr)
			if webErr != &WebApiErrorServerError {
				return c.SendStatus(401)
			} else {
				return c.SendStatus(500)
			}
		}

		newCommentId, err := app.GetCommentManager().CreateCommentForComment(*session, commentId, req.Text)

		if err != nil {
			return c.SendStatus(500)
		}

		type NewCommentResponse struct {
			Id string `json:"id"`
		}

		res := NewCommentResponse{fmt.Sprintf("%d", newCommentId)}

		return c.JSON(res)
	})

	group.Patch("/:forum/:postId/:commentId", func(c *fiber.Ctx) error {
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

		type PatchComment struct {
			Text *string `json:"text"`
		}

		var req PatchComment

		if err := c.BodyParser(&req); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ok, err := app.GetCommentManager().PatchComment(*session, commentId, req.Text)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if ok {
			return c.SendStatus(fiber.StatusNoContent)
		} else {
			return c.SendStatus(fiber.StatusBadRequest)
		}
	})

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

		return c.JSON(commentsToCommentsDto(comments))
	})
}
