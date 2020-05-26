package cmd

import (
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/dto"
)

type SaveCustomOAuthConfig struct {
	Config *models.CreateEditOAuthConfig
}

type ParseOAuthRawProfile struct {
	Provider string
	Body     string

	Result *dto.OAuthUserProfile
}
