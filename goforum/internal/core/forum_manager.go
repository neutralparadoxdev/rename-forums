package core

type ForumManager struct {
	db Database
}

func NewForumManager(db Database) *ForumManager {
	return &ForumManager{
		db: db,
	}
}

func (man *ForumManager) CreateForum(title, description string, ownerId int64) error {
	repo := man.db.GetForumRepository()
	return repo.Create(title, description, ownerId)
}