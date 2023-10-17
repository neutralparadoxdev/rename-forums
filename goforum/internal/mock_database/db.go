package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type MockDatabase struct {
	session SessionRepository
	user    UserRepository
}

func New() *MockDatabase {
	return &MockDatabase{
		session: SessionRepository{},
		user:    *NewUserRepository(),
	}
}

func (db *MockDatabase) Init() error {
	return nil
}

func (db *MockDatabase) GetSessionRepository() core.SessionRepository {
	return &db.session
}

func (db *MockDatabase) GetUserRepository() core.UserRepository {
	return &db.user
}
