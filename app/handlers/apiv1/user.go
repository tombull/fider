package apiv1

import (
	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/enum"
	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/errors"
	"github.com/tombull/teamdream/app/pkg/web"
)

// ListUsers returns all registered users
func ListUsers() web.HandlerFunc {
	return func(c *web.Context) error {
		allUsers := &query.GetAllUsers{}
		if err := bus.Dispatch(c, allUsers); err != nil {
			return c.Failure(err)
		}
		return c.Ok(allUsers.Result)
	}
}

// CreateUser is used to create new users
func CreateUser() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(actions.CreateUser)
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		var user *models.User

		getByReference := &query.GetUserByProvider{Provider: "reference", UID: input.Model.Reference}
		err := bus.Dispatch(c, getByReference)
		user = getByReference.Result

		if err != nil && errors.Cause(err) == app.ErrNotFound {
			if input.Model.Email != "" {
				getByEmail := &query.GetUserByEmail{Email: input.Model.Email}
				err = bus.Dispatch(c, getByEmail)
				user = getByEmail.Result
			}
			if err != nil && errors.Cause(err) == app.ErrNotFound {
				user = &models.User{
					Tenant: c.Tenant(),
					Name:   input.Model.Name,
					Email:  input.Model.Email,
					Role:   enum.RoleVisitor,
				}
				err = bus.Dispatch(c, &cmd.RegisterUser{User: user})
			}
		}

		if err != nil {
			return c.Failure(err)
		}

		if input.Model.Reference != "" && !user.HasProvider("reference") {
			if err := bus.Dispatch(c, &cmd.RegisterUserProvider{
				UserID:       user.ID,
				ProviderName: "reference",
				ProviderUID:  input.Model.Reference,
			}); err != nil {
				return c.Failure(err)
			}
		}

		return c.Ok(web.Map{
			"id": user.ID,
		})
	}
}
