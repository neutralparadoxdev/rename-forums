package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
	"log"
	"context"
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

	var forum core.Forum

	err := repo.db.pool.QueryRow(context.Background(), "SELECT * FROM forum WHERE title = $1", name).Scan(&forum.Title, &forum.Description, &forum.IsPublic)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		return nil, errors.New("GetByName:Query Row failed")
	}

	return &forum, nil
}

func (repo *ForumRepository) Delete(forum core.Forum) error {
	commandTag, err := repo.db.pool.Exec(context.Background(), "DELETE FROM forum WHERE title = $", forum.Title)

	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("No row deleted")
	}
	return nil
}

func (repo *ForumRepository) Create(title string, description string, ownerId int64, isPublic bool) error {
	return errors.New("Not Implemented")
}

func (repo *ForumRepository) GetAll(userId *int64) ([]core.Forum, error) {
	return nil, errors.New("Not Implemented")
}
