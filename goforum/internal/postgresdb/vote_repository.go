package postgresdb

import (
	"errors"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

type VoteRepository struct {
	db *PostgresDatabase
}

func NewVoteRepository(db *PostgresDatabase) core.VoteRepository {
	return &VoteRepository {
		db: db,
	}
}

func (repo *VoteRepository) HasVotedOn(postId int64, userId int64) (bool, error) {
	return false, errors.New("Not Implemented")
}

func (repo *VoteRepository) ChangeVote(postId int64, userId int64, vote int64) (int64, error) {
	return 0, errors.New("Not Implemented")
}

func (repo *VoteRepository) Vote(postId int64, userId int64, direction int64) (int64, error) {
	return 0, errors.New("Not Implemented")
}

func (repo *VoteRepository) GetVotesForPosts(userId int64, postIds []int64) ([]int64, error) {
	return nil, errors.New("Not Implemented")
}

