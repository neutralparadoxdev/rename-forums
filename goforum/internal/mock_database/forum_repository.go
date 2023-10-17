package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type ForumRepository struct {
	forums map[string]core.Forum
}

func NewForumRepository() *ForumRepository {
	forums := make(map[string]core.Forum)

	forums["Science"] = core.Forum{
		Title:        "Science",
		Description:  "Some Science",
		OwnerListIds: append(make([]int64, 1), 12),
	}

	forums["Math"] = core.Forum{
		Title:        "Math",
		Description:  "Math Forum",
		OwnerListIds: append(make([]int64, 1), 12),
	}

	return &ForumRepository{
		forums: forums,
	}
}

func (repo *ForumRepository) GetByName(name string) (*core.Forum, error) {
	return nil, nil
}

func (repo *ForumRepository) Delete(forum core.Forum) error {
	return nil
}

func (repo *ForumRepository) Create(title, description string, ownerId int64) error {
	repo.forums[title] = core.Forum{
		Description:  description,
		OwnerListIds: append(make([]int64, 1), ownerId),
	}
	return nil
}

func (repo *ForumRepository) GetAll() ([]core.Forum, error) {
	return make([]core.Forum, 0), nil
}
