package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type SessionRepository struct {
}

func (repo *SessionRepository) GetById() (*core.Session, error) {
	return nil, nil
}

func (repo *SessionRepository) Delete(session core.Session) error {
	return nil
}

func (repo *SessionRepository) Create(userid string, username string) (*core.Session, error) {
	return nil, nil
}

func (repo *SessionRepository) Save(session core.Session) error {
	return nil
}
