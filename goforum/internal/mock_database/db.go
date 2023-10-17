package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type MockDatabase struct {
	session SessionRepository
	user    UserRepository
	forum   ForumRepository
	post    PostRepository
}

func New() *MockDatabase {
	return &MockDatabase{
		session: SessionRepository{},
		user:    *NewUserRepository(),
		forum:   *NewForumRepository(),
		post:    *NewPostRepository(),
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

func (db *MockDatabase) GetForumRepository() core.ForumRepository {
	return &db.forum
}

func (db *MockDatabase) GetPostRepository() core.PostRepository {
	return &db.post
}
