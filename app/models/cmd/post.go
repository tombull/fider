package cmd

import (
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/enum"
)

type AddNewPost struct {
	Title       string
	Description string

	Result *models.Post
}

type UpdatePost struct {
	Post        *models.Post
	Title       string
	Description string

	Result *models.Post
}

type SetPostResponse struct {
	Post   *models.Post
	Text   string
	Status enum.PostStatus
}
