package core

type Forum struct {
	Title           string
	Description     string
	OwnerListIds    []int64
	UserJoinListIds []int64
	IsPublic        bool
}

func (f *Forum) IsOwner(userId int64) bool {
	for _, v := range f.OwnerListIds {
		if v == userId {
			return true
		}
	}
	return false
}

func (f *Forum) HasJoined(userId int64) bool {
	for _, v := range f.UserJoinListIds {
		if v == userId {
			return true
		}
	}
	return false
}

func (f *Forum) CanPost(userId int64) bool {
	return f.IsOwner(userId) || f.HasJoined(userId)
}

func (f *Forum) CanViewPosts(userId *int64) bool {

	if userId == nil {
		return f.IsPublic
	} else {
		return f.IsPublic || f.CanPost(*userId)
	}
}
