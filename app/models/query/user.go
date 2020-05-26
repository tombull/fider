package query

import "github.com/tombull/teamdream/app/models"

type CountUsers struct {
	Result int
}

type UserSubscribedTo struct {
	PostID int

	Result bool
}

type GetUserByAPIKey struct {
	APIKey string

	Result *models.User
}

type GetCurrentUserSettings struct {
	Result map[string]string
}

type GetUserByID struct {
	UserID int

	Result *models.User
}

type GetUserByEmail struct {
	Email string

	Result *models.User
}

type GetUserByProvider struct {
	Provider string
	UID      string

	Result *models.User
}

type GetAllUsers struct {
	Result []*models.User
}
