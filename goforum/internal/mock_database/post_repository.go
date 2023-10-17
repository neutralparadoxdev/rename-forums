package mockdatabase

import "github.com/neutralparadoxdev/rename-forums/goforum/internal/core"

type PostRepository struct {
	posts map[int64]core.Post
}

func NewPostRepository() *PostRepository {
	posts := make(map[int64]core.Post)

	posts[1] = core.Post{
		Title:           "Math first post",
		Body:            "This is the body for the post",
		Id:              1,
		ForumPostedName: "math",
		OwnerId:         12,
		AuthorName:      "hello",
	}
	return &PostRepository{
		posts: posts,
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
