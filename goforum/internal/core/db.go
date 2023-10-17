package core

type Database interface {
	GetSessionRepository() SessionRepository
	GetUserRepository() UserRepository
	GetForumRepository() ForumRepository
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
	GetAll() ([]Forum, error)
}

type UserRepository interface {
	GetById() (*User, error)
	GetByName(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(user User) error
	Create(user User) (*User, error)
	Save(user User) error
}
