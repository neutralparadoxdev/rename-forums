package core

import "time"

type Comment struct {
	Id int64

	Text string

	// subcomment Ids
	SubComments []int64

	// post that owns this comment
	PostOwner *int64

	/// comment that owns this comment
	CommentOwner *int64

	/// user owner
	Owner int64

	CreatedAt time.Time

	ModifiedAt time.Time
}
