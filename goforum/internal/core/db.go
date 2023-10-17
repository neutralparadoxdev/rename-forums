package core

type Database interface {
	GetSessionRepository() SessionRepository
	GetUserRepository() UserRepository
	GetForumRepository() ForumRepository
	GetPostRepository() PostRepository
	Init() error
}

type SessionRepository interface {
	GetById() (*Session, error)
	DoesSessionExist(session Session) (bool, error)
	Delete(session Session) error
	Create(userid string, username string) (*Session, error)
	Save(session Session) error
}

type ForumRepository interface {
	GetByName(name string) (*Forum, error)
	Delete(forum Forum) error
	Create(title string, description string, ownerId int64) error
	GetAll(userId *int64) ([]Forum, error)
}

type UserRepository interface {
	GetById(id int64) (*User, error)
	GetByName(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(user User) error
	Create(user User) (*User, error)
	Save(user User) error
}

type PostRepository interface {
	GetPostsOnForum(forumName string) ([]Post, error)
	Create(title string, body string, forumName string, userId int64) (bool, error)
	GetPost(userId int64) (*Post, error)
}
