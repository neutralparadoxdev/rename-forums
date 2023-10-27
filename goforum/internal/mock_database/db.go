package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type MockDatabase struct {
	session *SessionRepository
	user    *UserRepository
	forum   *ForumRepository
	post    *PostRepository
	vote    *VoteRepository
}

func New() *MockDatabase {

	user := NewUserRepository()
	return &MockDatabase{
		session: NewSessionRepository(),
		user:    user,
		forum:   NewForumRepository(),
		post:    NewPostRepository(user),
		vote:    NewVoteRepository(),
	}
}

func (db *MockDatabase) Init() error {
	return nil
}

func (db *MockDatabase) GetSessionRepository() core.SessionRepository {
	return db.session
}

func (db *MockDatabase) GetUserRepository() core.UserRepository {
	return db.user
}

func (db *MockDatabase) GetForumRepository() core.ForumRepository {
	return db.forum
}

func (db *MockDatabase) GetPostRepository() core.PostRepository {
	return db.post
}

func (db *MockDatabase) GetVoteRepository() core.VoteRepository {
	return db.vote
}
