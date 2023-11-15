package core

type CommentManager struct {
	database Database
}

func NewCommentManager(database Database) *CommentManager {
	return &CommentManager{
		database: database,
	}
}
