package core

type PostManager struct {
	db Database
}

func NewPostManager(db Database) *PostManager {
	return &PostManager{
		db: db,
	}
}

func (man *PostManager) GetPosts(forumName string) ([]Post, error) {
	repo := man.db.GetPostRepository()

	posts, err := repo.GetPostsOnForum(forumName)

	if err != nil {
		return make([]Post, 0), err
	}

	return posts, nil

}
