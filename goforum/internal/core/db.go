package core

type Database interface {
	GetSessionRepository() SessionRepository
	GetUserRepository() UserRepository
	GetForumRepository() ForumRepository
	GetPostRepository() PostRepository
	GetVoteRepository() VoteRepository
	GetCommentRepository() CommentRepository
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
	Create(title string, description string, ownerId int64, isPublic bool) error
	GetAll(userId *int64) ([]Forum, error)
}

type UserRepository interface {
	GetById(id int64) (*User, error)
	GetByName(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(user User) error
	Create(user User) (*User, error)
	Save(user User) error
	GetUserNamesForIds(ids []int64) map[int64]string
}

type PostRepository interface {
	GetPostsOnForum(forumName string) ([]Post, error)
	Create(title string, body string, forumName string, userId int64) (int64, error)
	Delete(userId int64, postId int64) (bool, error)
	GetPost(userId int64) (*Post, error)
	AddVote(postId int64, vote int64) error
	RemoveVote(postId int64, vote int64) error
	Patch(userId int64, postId int64, title *string, body *string) (bool, error)
}

type VoteRepository interface {
	HasVotedOn(postId int64, userId int64) (bool, error)
	/// returns the original vote. 0 means not voted
	ChangeVote(postId int64, userId int64, vote int64) (int64, error)
	Vote(postId int64, userId int64, direction int64) (int64, error)

	GetVotesForPosts(userId int64, postIds []int64) ([]int64, error)
}

type CommentRepository interface {
	DeleteComment(userId int64, commentId int64) (bool, error)

	/// Mark the comment as deleted
	MarkDeleted(userId int64, commentId int64) (bool, error)

	/// returns the new comment id or error
	NewComment(postId *int64, commentId *int64, userId int64, text string) (int64, error)

	/// Patch a comment
	PatchComment(userId int64, commentId int64, text string) (bool, error)

	/// returns one or more comments
	GetComment(commentId int64, depth int64) ([]Comment, error)

	/// returns one or more comments for post
	GetCommentForPost(postId int64, depth int64) ([]Comment, error)
}
