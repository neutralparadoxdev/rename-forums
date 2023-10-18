package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type SessionRepository struct {
	sessions map[string]string
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		sessions: make(map[string]string),
	}
}

func (repo *SessionRepository) DoesSessionExist(session core.Session) (bool, error) {
	_, exists := repo.sessions[session.Session]
	return exists, nil
}

func (repo *SessionRepository) Delete(session core.Session) error {
	delete(repo.sessions, session.Session)
	return nil
}

func (repo *SessionRepository) Save(session core.Session) error {
	repo.sessions[session.Session] = session.Session
	return nil
}
