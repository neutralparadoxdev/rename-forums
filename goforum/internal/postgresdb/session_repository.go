package postgresdb

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
	"errors"
)

type SessionRepository struct {
	db *PostgresDatabase
}

func NewSessionRepository(db *PostgresDatabase) core.SessionRepository {
	return &SessionRepository {
		db: db,
	}
}

func (repo *SessionRepository) DoesSessionExist(session core.Session) (bool, error) {
	return false, errors.New("Not Implemented") 
}

func (repo *SessionRepository) Delete(session core.Session) error {
	return errors.New("Not Implemented")
}

func (repo *SessionRepository) Save(session core.Session) error {
	return errors.New("Not Implemented")
}

