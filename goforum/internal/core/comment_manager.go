package core

type CommentManager struct {
	database Database
}

func NewCommentManager(database Database) *CommentManager {
	return &CommentManager{
		database: database,
	}
}

func (man *CommentManager) CreateCommentForPost(userId int64, postId int64, text string) (int64, error) {
	return man.database.GetCommentRepository().NewComment(&postId, nil, userId, text)
}

func (man *CommentManager) CreateCommentForComment(userId int64, commentId int64, text string) (int64, error) {
	return man.database.GetCommentRepository().NewComment(nil, &commentId, userId, text)
}

func (man *CommentManager) GetCommentWithUserSession(commentId int64, postId int64, forum string, session *Session) ([]Comment, error) {
	return man.database.GetCommentRepository().GetComment(commentId, 3)
}
