package handlers

import (
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/web"
)

// TotalUnreadNotifications returns the total number of unread notifications
func TotalUnreadNotifications() web.HandlerFunc {
	return func(c *web.Context) error {
		q := &query.CountUnreadNotifications{}
		if err := bus.Dispatch(c, q); err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{
			"total": q.Result,
		})
	}
}

// Notifications is the home for unread and recent notifications
func Notifications() web.HandlerFunc {
	return func(c *web.Context) error {
		q := &query.GetActiveNotifications{}
		if err := bus.Dispatch(c, q); err != nil {
			return c.Failure(err)
		}

		return c.Page(web.Props{
			Title:     "Notifications",
			ChunkName: "MyNotifications.page",
			Data: web.Map{
				"notifications": q.Result,
			},
		})
	}
}

// ReadNotification marks it as read and redirect to its content
func ReadNotification() web.HandlerFunc {
	return func(c *web.Context) error {
		id, err := c.ParamAsInt("id")
		if err != nil {
			return c.Failure(err)
		}

		q := &query.GetNotificationByID{ID: id}
		if err := bus.Dispatch(c, q); err != nil {
			return c.Failure(err)
		}

		if err = bus.Dispatch(c, &cmd.MarkNotificationAsRead{ID: q.Result.ID}); err != nil {
			return c.Failure(err)
		}

		return c.Redirect(c.BaseURL() + q.Result.Link)
	}
}

// ReadAllNotifications marks all unread notifications as read
func ReadAllNotifications() web.HandlerFunc {
	return func(c *web.Context) error {

		if err := bus.Dispatch(c, &cmd.MarkAllNotificationsAsRead{}); err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{})
	}
}
