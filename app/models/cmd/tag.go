package cmd

import "github.com/tombull/teamdream/app/models"

type AddNewTag struct {
	Name     string
	Color    string
	IsPublic bool

	Result *models.Tag
}

type UpdateTag struct {
	TagID    int
	Name     string
	Color    string
	IsPublic bool

	Result *models.Tag
}

type DeleteTag struct {
	Tag *models.Tag
}

type AssignTag struct {
	Tag  *models.Tag
	Post *models.Post
}

type UnassignTag struct {
	Tag  *models.Tag
	Post *models.Post
}
