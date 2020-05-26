package query

import "github.com/tombull/teamdream/app/models"

type GetAttachments struct {
	Post    *models.Post
	Comment *models.Comment

	Result []string
}
