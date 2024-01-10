package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
)

type ForumRepository struct {
	db *PostgresDatabase
}

func NewForumRepository(db *PostgresDatabase) core.ForumRepository {
	return &ForumRepository {
		db: db,
	}
}

func (repo *ForumRepository) GetByName(name string) (*core.Forum, error) {
	return nil, errors.New("Not implemented")
}

func (repo *ForumRepository) Delete(forum core.Forum) error {
	return errors.New("Not Implemented")
}

func (repo *ForumRepository) Create(title string, description string, ownerId int64, isPublic bool) error {
	return errors.New("Not Implemented")
}

func (repo *ForumRepository) GetAll(userId *int64) ([]core.Forum, error) {
	return nil, errors.New("Not Implemented")
}
