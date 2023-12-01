package webapi

import (
	"fmt"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type CommentDTO struct {
	Text         string `json:"text"`
	CommentOwner string `json:"commentOwner"`
	PostOwner    string `json:"postOwner"`
	Id           string `json:"id"`
	WasDeleted   bool   `json:"was_deleted"`
	Comments     []CommentDTO `json:"comments"`
}

func commentsToCommentsDto(comments []core.Comment) []CommentDTO {
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
			WasDeleted:   comments[i].WasDeleted,
			Comments:     commentsToCommentsDto(comments[i].SubComments),
		}
		commentsDto = append(commentsDto, commentDto)
	}

	return commentsDto
}
