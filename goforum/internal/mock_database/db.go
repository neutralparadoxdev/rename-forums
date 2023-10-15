package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type MockDatabase struct {
}

func new() MockDatabase {
	return MockDatabase{}
}

func (db *MockDatabase) Init() error {
	return nil
}

func (db *MockDatabase) GetSessionRepository() core.SessionRepository {
	return &SessionRepository{}
}

func (db *MockDatabase) GetUserRepository() core.UserRepository {
	return &UserRepository{}
}
