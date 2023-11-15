package core

type Comment struct {
	Id int64

	Text string

	// subcomment Ids
	SubComments []int64

	// post that owns this comment
	PostOwner *int64

	/// comment that owns this comment
	CommentOwner *int64
}
