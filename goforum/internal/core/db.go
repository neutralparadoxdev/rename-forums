package core

type Database interface {
	GetSessionRepository() SessionRepository
	GetUserRepository() UserRepository
	GetForumRepository() ForumRepository
	GetPostRepository() PostRepository
	GetVoteRepository() VoteRepository
	Init() error
}

type SessionRepository interface {
	DoesSessionExist(session Session) (bool, error)
	Delete(session Session) error
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
	Create(title string, body string, forumName string, userId int64) (int64, error)
	GetPost(userId int64) (*Post, error)
	AddVote(postId int64, vote int64) error
	RemoveVote(postId int64, vote int64) error
}

type VoteRepository interface {
	HasVotedOn(postId int64, userId int64) (bool, error)
	/// returns the original vote. 0 means not voted
	ChangeVote(postId int64, userId int64, vote int64) (int64, error)
	Vote(postId int64, userId int64, direction int64) (int64, error)

	GetVotesForPosts(userId int64, postIds []int64) ([]int64, error)
}
