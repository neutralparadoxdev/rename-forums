package core

import "time"

type Post struct {
	Title           string
	Body            string
	Id              int64
	OwnerId         int64
	ForumPostedName string
	CreatedAt       time.Time
	LastEdited      time.Time
	AuthorName      string
}
