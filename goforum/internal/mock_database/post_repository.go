package mockdatabase

import (
	"errors"
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

	posts[2] = core.Post{
		Title: "Dev Markdown Example",
		Body: `
Some Demonstration Markdown

# h1

## h2

### h3

#### h4

##### h5

###### h6

heading 1
=========

heading 2
---------

> This is a block
> quote.

1. List 1
2. List 2
3. List 3
4. List 4
5. List
6. List
7. List
8. List
9. List
10. List

- Another list
- Together with a numbered list

A paragraph

A paragraph with *emphasis*

Some more **emphasis**

A [Link](http://google.com)

Line ~~break~~
`,
		Id:              2,
		ForumPostedName: "math",
		OwnerId:         12,
		AuthorName:      "hello",
		CreatedAt:       time.Now(),
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

func (repo *PostRepository) Create(title string, body string, forumName string, userId int64) (int64, error) {
	id := rand.Int63()

	user, err := repo.users.GetById(userId)

	if err != nil {
		return 0, err
	}

	repo.posts[id] = core.Post{
		Title:           title,
		Body:            body,
		ForumPostedName: forumName,
		CreatedAt:       time.Now(),
		OwnerId:         userId,
		AuthorName:      user.Username,
		Id:              id,
	}
	return id, nil
}

func (repo *PostRepository) GetPost(postId int64) (*core.Post, error) {
	for _, v := range repo.posts {
		if v.Id == postId {
			return &v, nil
		}
	}
	return nil, nil
}

func (repo *PostRepository) AddVote(postId int64, vote int64) error {
	post, exists := repo.posts[postId]
	if exists {
		if vote == -1 {
			post.DownVote += 1
		}

		if vote == 1 {
			post.UpVote += 1
		}

		return nil
	} else {
		return errors.New("not_found")
	}
}

func (repo *PostRepository) RemoveVote(postId int64, vote int64) error {
	post, exists := repo.posts[postId]
	if exists {
		if vote == -1 {
			post.DownVote -= 1
		}

		if vote == 1 {
			post.UpVote -= 1
		}

		return nil
	} else {
		return errors.New("not_found")
	}

}

func (repo *PostRepository) Patch(userId int64, postId int64, title *string, body *string) (bool, error) {
	// check and validate post
	post, exists := repo.posts[postId]

	if exists {
		if post.OwnerId == userId {
			// patch
			if title != nil {
				// patch user
				post.Title = *title
			}

			if body != nil {
				// patch body
				post.Body = *body
			}

			repo.posts[postId] = post
			return true, nil
		} else {
			// owner is not correct
			return false, nil

		}
	} else {
		// post doesnt exist
		return false, nil
	}
}

func (repo *PostRepository) Delete(userId int64, postId int64) (bool, error) {
	// check and validate post
	post, exists := repo.posts[postId]

	if exists {
		if post.OwnerId == userId {
			delete(repo.posts, postId)
		} else {
			// owner is not correct
			return false, errors.New("owner_id_user_id_mismatch")
		}
	} else {
		// post doesnt exist
		return false, errors.New("post_not_found")
	}

	return true, nil
}
