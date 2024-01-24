package mockdatabase

import (
	"math/rand"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type CommentRepository struct {
	comments map[int64]core.Comment
}

func NewCommentRepository() *CommentRepository {

	comments := make(map[int64]core.Comment)

	num10ptr := new(int64)
	*num10ptr = 10
	comments[100] = core.Comment{
		PostOwner: num10ptr,
		Text:      "This is a comment for Math",
		Owner:     12,
		Id:        100,
	}

	num100ptr := new(int64)
	*num100ptr = 100

	comments[101] = core.Comment{
		CommentOwner: num100ptr,
		Text:         "This is a comment for Math",
		Owner:        12,
		Id:           101,
	}

	comments[102] = core.Comment{
		CommentOwner: num100ptr,
		Text:         "This is a comment for Math 2",
		Owner:        12,
		Id:           102,
	}

	return &CommentRepository{
		comments: comments,
	}
}

func (repo *CommentRepository) DeleteComment(userId int64, commentId int64) (bool, error) {
	comment, exists := repo.comments[commentId]
	if exists && comment.Owner == userId {
		delete(repo.comments, commentId)
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *CommentRepository) NewComment(postId *int64, commentId *int64, userId int64, text string) (int64, error) {
	id := rand.Int63()

	for id == 0 {
		id = rand.Int63()
	}

	newComment := core.Comment{
		PostOwner:    postId,
		CommentOwner: commentId,
		Owner:        userId,
		Text:         text,
		Id:           id,
	}
	repo.comments[id] = newComment

	return id, nil
}

func (repo *CommentRepository) PatchComment(userId int64, commentId int64, text string) (bool, error) {
	comment, exists := repo.comments[commentId]
	if exists && comment.Owner == userId {
		comment.Text = text
		repo.comments[commentId] = comment
		return true, nil
	} else {
		return false, nil
	}
}

/**
 * Retrieves the comment list for the provided comment. Recursively does the same for the returned
 * comments until depth is 0.
*/
func GetComments(id int64, depth int64, comments map[int64]core.Comment) []core.Comment {
	if depth == 0 {
		return make([]core.Comment, 0)
	}
	out := make([]core.Comment, 0)
	for _, val := range comments {
		if val.CommentOwner != nil && *val.CommentOwner == id {
			if depth > 0 {
				val.SubComments = GetComments(val.Id, depth-1, comments)
			}
			out = append(out, val)
		}
	}
	return out
}

func (repo *CommentRepository) GetComment(commentId int64, depth int64) ([]core.Comment, error) {
	val, exists := repo.comments[commentId]
	if exists {
		commentsOut := make([]core.Comment, 0)
		commentsOut = append(commentsOut, val)
		val.SubComments = GetComments(commentId, depth, repo.comments)
		return commentsOut, nil
	}

	return make([]core.Comment, 0), nil
}

func (repo *CommentRepository) GetCommentForPost(postId int64, depth int64) ([]core.Comment, error) {
	commentsOut := make([]core.Comment, 0)

	for _, val := range repo.comments {
		if val.PostOwner != nil && *val.PostOwner == postId {
			commentsOut = append(commentsOut, val)
		}
	}
	for i := range commentsOut {
		commentsOut[i].SubComments = GetComments(commentsOut[i].Id, 3, repo.comments)
	}

	return commentsOut, nil
}

func (repo *CommentRepository) MarkDeleted(userId int64, commentId int64) (bool, error) {
	comment, exists := repo.comments[commentId]
	if exists && comment.Owner == userId {
		comment.WasDeleted = true
		repo.comments[commentId] = comment
		return true, nil
	} else {
		return false, nil
	}
}
