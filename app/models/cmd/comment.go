package cmd

import "github.com/tombull/teamdream/app/models"

type AddNewComment struct {
	Post    *models.Post
	Content string

	Result *models.Comment
}

type UpdateComment struct {
	CommentID int
	Content   string
}

type DeleteComment struct {
	CommentID int
}
