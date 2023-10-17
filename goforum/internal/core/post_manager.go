package core

import "errors"

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

func (man *PostManager) GetPost(id int64, forumName string, userId *int64) (*Post, error) {
	repoForum := man.db.GetForumRepository()

	forum, err := repoForum.GetByName(forumName)

	if err != nil {
		return nil, errors.New("get_post: forum error")
	}

	if forum == nil {
		return nil, errors.New("get_post: forum not found")
	}

	if !forum.CanViewPosts(userId) {
		return nil, errors.New("get_post: User does not have permission to post")
	}

	repoPost := man.db.GetPostRepository()

	post, err := repoPost.GetPost(id)

	if err != nil {
		return nil, errors.New("get_post: Could not retrieve post")
	}

	if post == nil {
		return nil, nil
	}

	return post, nil
}
