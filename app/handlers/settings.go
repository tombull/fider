package handlers

import (
	"time"

	"github.com/tombull/teamdream/app/models/enum"
	"github.com/tombull/teamdream/app/models/query"

	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/pkg/bus"

	"github.com/tombull/teamdream/app/tasks"

	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/pkg/web"
)

// ChangeUserEmail register the intent of changing user email
func ChangeUserEmail() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(actions.ChangeUserEmail)
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		err := bus.Dispatch(c, &cmd.SaveVerificationKey{
			Key:      input.Model.VerificationKey,
			Duration: 24 * time.Hour,
			Request:  input.Model,
		})
		if err != nil {
			return c.Failure(err)
		}

		c.Enqueue(tasks.SendChangeEmailConfirmation(input.Model))

		return c.Ok(web.Map{})
	}
}

// VerifyChangeEmailKey checks if key is correct and update user's email
func VerifyChangeEmailKey() web.HandlerFunc {
	return func(c *web.Context) error {
		result, err := validateKey(enum.EmailVerificationKindChangeEmail, c)
		if result == nil {
			return err
		}

		if result.UserID != c.User().ID {
			return c.Redirect(c.BaseURL())
		}

		changeEmail := &cmd.ChangeUserEmail{
			UserID: result.UserID,
			Email:  result.Email,
		}
		if err = bus.Dispatch(c, changeEmail); err != nil {
			return c.Failure(err)
		}

		err = bus.Dispatch(c, &cmd.SetKeyAsVerified{Key: result.Key})
		if err != nil {
			return c.Failure(err)
		}
		return c.Redirect(c.BaseURL() + "/settings")
	}
}

// UserSettings is the current user's profile settings page
func UserSettings() web.HandlerFunc {
	return func(c *web.Context) error {
		settings := &query.GetCurrentUserSettings{}
		if err := bus.Dispatch(c, settings); err != nil {
			return err
		}

		return c.Page(web.Props{
			Title:     "Settings",
			ChunkName: "MySettings.page",
			Data: web.Map{
				"userSettings": settings.Result,
			},
		})
	}
}

// UpdateUserSettings updates current user settings
func UpdateUserSettings() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(actions.UpdateUserSettings)
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		if err := bus.Dispatch(c,
			&cmd.UploadImage{
				Image:  input.Model.Avatar,
				Folder: "avatars",
			},
			&cmd.UpdateCurrentUser{
				Name:       input.Model.Name,
				Avatar:     input.Model.Avatar,
				AvatarType: input.Model.AvatarType,
			},
			&cmd.UpdateCurrentUserSettings{
				Settings: input.Model.Settings,
			},
		); err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{})
	}
}

// ChangeUserRole changes given user role
func ChangeUserRole() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(actions.ChangeUserRole)
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		changeRole := &cmd.ChangeUserRole{
			UserID: input.Model.UserID,
			Role:   input.Model.Role,
		}

		if err := bus.Dispatch(c, changeRole); err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{})
	}
}

// DeleteUser erases current user personal data and sign them out
func DeleteUser() web.HandlerFunc {
	return func(c *web.Context) error {
		if err := bus.Dispatch(c, &cmd.DeleteCurrentUser{}); err != nil {
			return c.Failure(err)
		}

		c.RemoveCookie(web.CookieAuthName)
		return c.Ok(web.Map{})
	}
}

// RegenerateAPIKey regenerates current user's API Key
func RegenerateAPIKey() web.HandlerFunc {
	return func(c *web.Context) error {
		regenerateAPIKey := &cmd.RegenerateAPIKey{}
		if err := bus.Dispatch(c, regenerateAPIKey); err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{
			"apiKey": regenerateAPIKey.Result,
		})
	}
}
