package query

import "github.com/tombull/teamdream/app/models"

type ListPostVotes struct {
	PostID int
	Limit  int

	Result []*models.Vote
}
