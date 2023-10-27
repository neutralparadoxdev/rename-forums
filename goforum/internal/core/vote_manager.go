package core

type VoteManager struct {
	db Database
}

func NewVoteManager(db Database) *VoteManager {
	return &VoteManager{
		db: db,
	}
}

func (man *VoteManager) HasVotedOn(userId int64, postId int64) (bool, error) {
	return man.db.GetVoteRepository().HasVotedOn(postId, userId)
}

// returns the previous direction voted. 0 means not voted
func (man *VoteManager) ChangeVote(userId int64, postId int64, vote int64) (int64, error) {
	return man.db.GetVoteRepository().ChangeVote(postId, userId, vote)
}

// returns the votes for the given posts in order. 1 upvote,  -1 downvote, 0 not voted
func (man *VoteManager) GetVotesForPosts(userId int64, postIds []int64) ([]int64, error) {
	return man.db.GetVoteRepository().GetVotesForPosts(userId, postIds)

}
