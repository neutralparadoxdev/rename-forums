package core

import "time"

type CommentBuilder struct { 
	Comment
}

func NewCommentBuilder() CommentBuilder {
	return CommentBuilder {
		Comment: NewComment("", 0),
	}
}

func (b *CommentBuilder) AddId(id int64)  *CommentBuilder {
	b.Comment.Id = id
	return b
}

func (b *CommentBuilder) AddText(text string)  *CommentBuilder {
	b.Comment.Text = text
	return  b
}

func (b *CommentBuilder) AddSubComment(comment Comment)  *CommentBuilder {
	if b.Comment.SubComments == nil {
		b.Comment.SubComments = make([]Comment, 1)
		b.Comment.SubComments[0] = comment
	}
	b.Comment.SubComments = append(b.Comment.SubComments, comment)
	return b
}

func (b *CommentBuilder) AddOwner(owner int64) *CommentBuilder {
	b.Comment.Owner = owner
	return b
}

func (b *CommentBuilder) AddCommentOwner(owner int64) *CommentBuilder {
	b.Comment.CommentOwner = &owner
	return b
}

func (b *CommentBuilder) AddPostOwner(owner int64) *CommentBuilder {
	b.Comment.PostOwner = &owner
	return b
}

func (b *CommentBuilder) Make() Comment {
	return b.Comment
}

type Comment struct {
	Id int64

	Text string

	// comments under this comment
	SubComments []Comment

	// post that owns this comment
	PostOwner *int64

	/// comment that owns this comment
	CommentOwner *int64

	/// user owner
	Owner int64

	CreatedAt time.Time

	ModifiedAt time.Time

	IsHidden bool

	WasDeleted bool

	MoreCommentsAvailable func() bool
}

func NewComment(
	text string,
	owner int64,
) Comment {
	return Comment {
		Text: text,
		Owner: owner,
		SubComments: make([]Comment, 0),
	}
}


/// Returns whether there is more comments available that provided in object
func (comment *Comment) IsMoreCommentsAvailable() bool {
	if comment.MoreCommentsAvailable != nil {
		return comment.MoreCommentsAvailable()
	}
	return false
}
