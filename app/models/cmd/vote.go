package cmd

import "github.com/tombull/teamdream/app/models"

type AddVote struct {
	Post *models.Post
	User *models.User
}

type RemoveVote struct {
	Post *models.Post
	User *models.User
}

type MarkPostAsDuplicate struct {
	Post     *models.Post
	Original *models.Post
}
