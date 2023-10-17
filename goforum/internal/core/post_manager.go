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

func (man *PostManager) CreatePost(title string, body string, forumName string, userId int64) (bool, error) {
	return man.db.GetPostRepository().Create(title, body, forumName, userId)
}
