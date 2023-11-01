package core

type ForumManager struct {
	db Database
}

func NewForumManager(db Database) *ForumManager {
	return &ForumManager{
		db: db,
	}
}

func (man *ForumManager) CreateForum(title, description string, ownerId int64, isPublic bool) error {
	repo := man.db.GetForumRepository()
	return repo.Create(title, description, ownerId, isPublic)
}

func (man *ForumManager) GetAll(userId *int64) ([]Forum, error) {
	repo := man.db.GetForumRepository()
	return repo.GetAll(userId)
}

func (man *ForumManager) GetForum(forumName string) (*Forum, error) {
	return man.db.GetForumRepository().GetByName(forumName)
}
