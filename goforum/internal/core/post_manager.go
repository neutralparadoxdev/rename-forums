package core

import (
	"errors"
	"sort"
	"log"
)

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

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.Before(posts[j].CreatedAt)
	})

	return posts, nil
}

func (man *PostManager) CreatePost(title string, body string, forumName string, userId int64) (int64, error) {
	return man.db.GetPostRepository().Create(title, body, forumName, userId)
}

func (man *PostManager) GetPost(id int64, forumName string, userId *int64, includeComments bool) (*Post, error) {
	repoForum := man.db.GetForumRepository()

	forum, err := repoForum.GetByName(forumName)

	if err != nil {
		return nil, errors.New("get_post: forum error")
	}

	if forum == nil {
		return nil, errors.New("get_post: forum not found")
	}

	if !forum.CanViewPosts(userId) {
		return nil, errors.New("user_cant_post")
	}

	repoPost := man.db.GetPostRepository()

	post, err := repoPost.GetPost(id)

	if err != nil {
		return nil, errors.New("get_post: Could not retrieve post")
	}

	if post == nil {
		return nil, nil
	}

	if includeComments {
		comments, err := man.db.GetCommentRepository().GetCommentForPost(id, 3)

		if err != nil {
			return nil, errors.New("get_post: Could not retrieve comments")
		}

		userSet := make(map[int64]bool)

		commentStack := make([]*Comment, 0)

		for i := range comments {
			commentStack = append(commentStack, &comments[i])
		}

		for len(commentStack) > 0 {
			log.Printf("comment stack size: %d", len(commentStack))
			comment := commentStack[len(commentStack)-1]

			commentStack = commentStack[0:len(commentStack)-1]
	
			userSet[comment.Owner] = true
	
			if comment.SubComments == nil {
				continue
			}
	
			for i := range comment.SubComments {
				commentStack = append(commentStack, &comment.SubComments[i])
			}
		}
	
		log.Printf("USER IDS")
		for k, _ := range userSet {
			log.Printf("%d\n", k)
		}

		ids := make([]int64, 0, 10)

		for k, _ := range userSet {
			ids = append(ids, k)
		}

		usernames := man.db.GetUserRepository().GetUserNamesForIds(ids)

		for k, v := range usernames {
			log.Printf("%d, %s", k, v)
		}


		commentStack = make([]*Comment, 0)

		for i := range comments {
			commentStack = append(commentStack, &comments[i])
		}

		for len(commentStack) > 0 {
			comment := commentStack[len(commentStack)-1]


			copy := usernames[comment.Owner]

			comment.Username = &copy

			commentStack = commentStack[0:len(commentStack)-1]
	
			if comment.SubComments == nil {
				continue
			}
	
			for i := range comment.SubComments {
				commentStack = append(commentStack, &comment.SubComments[i])
			}
		}

		post.Comments = &comments
	}

	return post, nil
}

func (man *PostManager) GetAllPosts(userId *int64) ([]Post, error) {
	if userId != nil {
		forums, err := man.db.GetForumRepository().GetAll(userId)

		if err != nil {
			return make([]Post, 0), err
		}

		out := make([]Post, 0)

		for _, v := range forums {
			posts, err := man.db.GetPostRepository().GetPostsOnForum(v.Title)
			if err != nil {
				return make([]Post, 0), err
			}

			for _, i := range posts {
				i.ForumPostedName = v.Title
				out = append(out, i)
			}
		}
		return out, nil
	} else {
		forums, err := man.db.GetForumRepository().GetAll(userId)

		if err != nil {
			return make([]Post, 0), err
		}

		out := make([]Post, 0)

		for _, v := range forums {
			posts, err := man.db.GetPostRepository().GetPostsOnForum(v.Title)
			if err != nil {
				return make([]Post, 0), err
			}

			for _, i := range posts {
				out = append(out, i)
			}
		}
		return out, nil

	}
}

func (man *PostManager) PatchPost(userId int64, postId int64, title *string, body *string) (bool, error) {
	return man.db.GetPostRepository().Patch(userId, postId, title, body)
}

func (man *PostManager) DeletePost(userId int64, postId int64) (bool, error) {
	return man.db.GetPostRepository().Delete(userId, postId)
}
