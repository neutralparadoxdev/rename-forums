package mockdatabase

import (
	"errors"

	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type postId = int64

type userId = int64

type VoteRepository struct {
	votes map[userId]map[postId]core.Vote
}

func NewVoteRepository() *VoteRepository {
	return &VoteRepository{}
}

func (repo *VoteRepository) HasVotedOn(postId int64, userId int64) (bool, error) {
	posts, exists := repo.votes[userId]

	if exists {
		_, exists := posts[postId]

		if exists {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, nil
	}
}

func (repo *VoteRepository) ChangeVote(postId int64, userId int64, vote int64) (int64, error) {
	posts, exists := repo.votes[userId]

	if exists {
		oldVote, exists := posts[postId]

		if exists {
			if vote == 0 {
				delete(posts, postId)
			} else {
				oldVote.Direction = vote
				posts[postId] = oldVote
			}

			return oldVote.Direction, nil
		} else {
			return 0, errors.New("not_found")
		}
	} else {
		return 0, errors.New("not_found")
	}

}

func (repo *VoteRepository) Vote(postId int64, userId int64, direction int64) (int64, error) {
	posts, exists := repo.votes[userId]

	if exists {
		_, exists := posts[postId]
		if !exists {
			posts[postId] = core.Vote{
				UserId:    userId,
				PostId:    postId,
				Direction: direction,
			}
			return 1, nil

		} else {
			return 0, errors.New("found_existing_vote")
		}

	} else {
		return 0, errors.New("not_found")
	}
}
