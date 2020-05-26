package actions

import (
	"context"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/pkg/validate"
)

// Actionable is any action that the user can perform using the web app
type Actionable interface {
	Initialize() interface{}
	IsAuthorized(ctx context.Context, user *models.User) bool
	Validate(ctx context.Context, user *models.User) *validate.Result
}
