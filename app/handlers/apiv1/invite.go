package apiv1

import (
	"strings"

	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/log"
	"github.com/tombull/teamdream/app/pkg/markdown"
	"github.com/tombull/teamdream/app/pkg/web"
	"github.com/tombull/teamdream/app/tasks"
)

// SendSampleInvite to current user's email
func SendSampleInvite() web.HandlerFunc {
	return func(c *web.Context) error {
		input := &actions.InviteUsers{IsSampleInvite: true}
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		if c.User().Email != "" {
			input.Model.Message = strings.Replace(input.Model.Message, app.InvitePlaceholder, "*the link to accept invitation will be here*", -1)
			to := dto.NewRecipient(c.User().Name, c.User().Email, dto.Props{
				"subject": input.Model.Subject,
				"message": markdown.Full(input.Model.Message),
			})

			bus.Publish(c, &cmd.SendMail{
				From:         c.Tenant().Name,
				To:           []dto.Recipient{to},
				TemplateName: "invite_email",
				Props: dto.Props{
					"logo": web.LogoURL(c),
				},
			})
		}

		return c.Ok(web.Map{})
	}
}

// SendInvites sends an email to each recipient
func SendInvites() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(actions.InviteUsers)
		if result := c.BindTo(input); !result.Ok {
			return c.HandleValidation(result)
		}

		log.Warnf(c, "Sending @{TotalInvites:magenta} invites by @{ClientIP:magenta}", dto.Props{
			"TotalInvites": len(input.Invitations),
			"ClientIP":     c.Request.ClientIP,
		})
		c.Enqueue(tasks.SendInvites(input.Model.Subject, input.Model.Message, input.Invitations))

		return c.Ok(web.Map{})
	}
}
