package mockdatabase

import (
	"math/rand"
	"time"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type PostRepository struct {
	posts map[int64]core.Post
	users *UserRepository
}

func NewPostRepository(users *UserRepository) *PostRepository {
	posts := make(map[int64]core.Post)

	posts[1] = core.Post{
		Title:           "Math first post",
		Body:            "This is the body for the post",
		Id:              1,
		ForumPostedName: "math",
		OwnerId:         12,
		AuthorName:      "hello",
		CreatedAt:       time.Now().Add(-1 * time.Minute),
	}

	posts[10] = core.Post{
		Title:           "Math second post",
		Body:            "This is the body for the post",
		Id:              10,
		ForumPostedName: "math",
		OwnerId:         12,
		AuthorName:      "hello",
		CreatedAt:       time.Now().Add(-2 * time.Minute),
	}

	posts[5] = core.Post{
		Title:           "Private Math first post",
		Body:            "This is the body for the post for privat emath",
		Id:              5,
		ForumPostedName: "privatemath",
		OwnerId:         12,
		AuthorName:      "hello",
	}

	return &PostRepository{
		posts: posts,
		users: users,
	}
}

func (repo *PostRepository) GetPostsOnForum(forumName string) ([]core.Post, error) {
	out := make([]core.Post, 0)

	for _, v := range repo.posts {
		if v.ForumPostedName == forumName {
			out = append(out, v)
		}
	}

	return out, nil
}

func (repo *PostRepository) Create(title string, body string, forumName string, userId int64) (bool, error) {
	id := rand.Int63()

	user, err := repo.users.GetById(userId)

	if err != nil {
		return false, err
	}

	repo.posts[id] = core.Post{
		Title:           title,
		Body:            body,
		ForumPostedName: forumName,
		CreatedAt:       time.Now(),
		OwnerId:         userId,
		AuthorName:      user.Username,
	}
	return true, nil
}

func (repo *PostRepository) GetPost(postId int64) (*core.Post, error) {
	for _, v := range repo.posts {
		if v.Id == postId {
			return &v, nil
		}
	}
	return nil, nil
}
