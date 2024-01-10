package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
)

type PostRepository struct {
	db *PostgresDatabase
}

func NewPostRepository(db *PostgresDatabase) core.PostRepository {
	return &PostRepository {
		db: db,
	}
}

func (repo *PostRepository) GetPostsOnForum(forumName string) ([]core.Post, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *PostRepository) Create(title string, body string, forumName string, userId int64) (int64, error) {
	return 0, errors.New("Not Implemented")
}


func (repo *PostRepository) Delete(userId int64, postId int64) (bool, error) {
	return false, errors.New("Not Implemented")
}

func (repo *PostRepository) GetPost(userId int64) (*core.Post, error) {
	return nil, errors.New("Not Implemented")
}

func (repo *PostRepository) AddVote(postId int64, vote int64) error {
	return errors.New("Not Implemented")
}

func (repo *PostRepository) RemoveVote(postId int64, vote int64) error {
	return errors.New("Not Implemented")
}

func (repo *PostRepository) Patch(userId int64, postId int64, title *string, body *string) (bool, error) {
	return false, errors.New("Not Implemented")
}

