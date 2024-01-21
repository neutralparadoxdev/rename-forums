package core

import ( 
	"errors"
	"log"
)

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
	comments, err := man.database.GetCommentRepository().GetComment(commentId, 3)

	if err != nil {
		return nil, err
	}

	userSet := make(map[int64]bool)

	commentStack := make([]*Comment, 0)
	
	for len(commentStack) > 0 {
		comment := commentStack[len(commentStack)-1]

		userSet[comment.Owner] = true

		if comment.SubComments == nil {
			continue
		}

		for i := range comment.SubComments {
			commentStack = append(commentStack, &comment.SubComments[i])
		}
	}
	
	for k, _ := range userSet {
		log.Printf("%d\n", k)
	}


	return comments, nil
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

func (man *CommentManager) DeleteComment(session Session, commentId int64) (bool, error) {
	userId, err := session.GetUserId()

	if err != nil {
		return false, errors.New("delete_comment: session_userid_error")
	}

	return man.database.GetCommentRepository().MarkDeleted(userId, commentId)
}
