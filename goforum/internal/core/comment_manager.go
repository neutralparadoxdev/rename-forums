package core

import "errors"

type CommentManager struct {
	database Database
}

func NewCommentManager(database Database) *CommentManager {
	return &CommentManager{
		database: database,
	}
}

func (man *CommentManager) CreateCommentForPost(session Session, postId int64, text string) (int64, error) {
	userId, err := session.GetUserId()

	if err != nil {
		return 0, errors.New("create_comment_for_posts: session_userid_error")
	}

	return man.database.GetCommentRepository().NewComment(&postId, nil, userId, text)
}

func (man *CommentManager) CreateCommentForComment(session Session, commentId int64, text string) (int64, error) {
	userId, err := session.GetUserId()

	if err != nil {
		return 0, errors.New("create_comment_for_comments: session_userid_error")
	}

	return man.database.GetCommentRepository().NewComment(nil, &commentId, userId, text)
}

func (man *CommentManager) GetCommentWithUserSession(commentId int64, postId int64, forum string, session *Session) ([]Comment, error) {
	return man.database.GetCommentRepository().GetComment(commentId, 3)
}

func (man *CommentManager) PatchComment(session Session, commentId int64, text *string) (bool, error) {
	userId, err := session.GetUserId()

	if err != nil {
		return false, errors.New("patch_comment: session_userid_error")
	}

	if text == nil {
		return true, nil
	}

	return man.database.GetCommentRepository().PatchComment(userId, commentId, *text)
}
