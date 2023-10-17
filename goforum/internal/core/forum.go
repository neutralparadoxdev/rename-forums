package core

type Forum struct {
	Title           string
	Description     string
	OwnerListIds    []int64
	UserJoinListIds []int64
	IsPublic        bool
}
