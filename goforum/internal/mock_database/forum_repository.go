package mockdatabase

import (
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type ForumRepository struct {
	forums map[string]core.Forum
}

func NewForumRepository() *ForumRepository {
	forums := make(map[string]core.Forum)

	forums["science"] = core.Forum{
		Title:           "science",
		Description:     "Some Science",
		IsPublic:        true,
		OwnerListIds:    append(make([]int64, 1), 12),
		UserJoinListIds: make([]int64, 0),
	}

	forums["math"] = core.Forum{
		Title:           "math",
		Description:     "Math Forum",
		IsPublic:        true,
		OwnerListIds:    append(make([]int64, 1), 12),
		UserJoinListIds: make([]int64, 0),
	}

	forums["privatemath"] = core.Forum{
		Title:           "privatemath",
		Description:     "Private Math Forum",
		IsPublic:        false,
		OwnerListIds:    append(make([]int64, 1), 12),
		UserJoinListIds: make([]int64, 0),
	}

	return &ForumRepository{
		forums: forums,
	}
}

func (repo *ForumRepository) GetByName(name string) (*core.Forum, error) {
	value, exists := repo.forums[name]

	if exists {
		return &value, nil
	} else {
		return nil, nil
	}
}

func (repo *ForumRepository) Delete(forum core.Forum) error {
	return nil
}

func (repo *ForumRepository) Create(title, description string, ownerId int64, isPublic bool) error {
	repo.forums[title] = core.Forum{
		Title:        title,
		Description:  description,
		OwnerListIds: append(make([]int64, 1), ownerId),
		IsPublic:     isPublic,
	}
	return nil
}

func (repo *ForumRepository) GetAll(userId *int64) ([]core.Forum, error) {
	out := make([]core.Forum, 0)

	if userId == nil {
		for _, forum := range repo.forums {
			if forum.IsPublic {
				out = append(out, forum)
				continue
			}
		}
		return out, nil
	}

	for _, forum := range repo.forums {
		if forum.IsPublic {
			out = append(out, forum)
			continue
		}

		if containsI64(&forum.OwnerListIds, *userId) {
			out = append(out, forum)
			continue
		}

		if containsI64(&forum.UserJoinListIds, *userId) {
			out = append(out, forum)
			continue
		}
	}

	return out, nil
}
