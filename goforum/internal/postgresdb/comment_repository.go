package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
)

type CommentRepository struct {
	db *PostgresDatabase
}

func NewCommentRepository(db *PostgresDatabase) core.CommentRepository {
	return &CommentRepository {
		db: db,
	}
}

func (repo *CommentRepository) DeleteComment(userId int64, commentId int64) (bool, error) {
	return false, errors.New("Not Implemented")
}

func (repo *CommentRepository) NewComment(postId *int64, commentId *int64, userId int64, text string) (int64, error) {
	return 0, errors.New("Not Implemented")
}

func (repo *CommentRepository) PatchComment(userId int64, commentId int64, text string) (bool, error) {
	return false, errors.New("Not Implemented")
}

func (repo *CommentRepository) GetComment(commentId int64, depth int64) ([]core.Comment, error) {
	return make([]core.Comment, 0), errors.New("Not Implemented")
}

func (repo *CommentRepository) GetCommentForPost(postId int64, depth int64) ([]core.Comment, error) {
	return make([]core.Comment, 0), errors.New("Not Implemented")
}

func (repo *CommentRepository) MarkDeleted(userId int64, commentId int64) (bool, error) {
	return false, errors.New("Not Implemented")
}

