package query

import "github.com/tombull/teamdream/app/models"

type GetCommentByID struct {
	CommentID int

	Result *models.Comment
}

type GetCommentsByPost struct {
	Post *models.Post

	Result []*models.Comment
}
