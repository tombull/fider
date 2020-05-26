package handlers

import (
	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/web"
)

// ManageTags is the home page for managing tags
func ManageTags() web.HandlerFunc {
	return func(c *web.Context) error {
		getAllTags := &query.GetAllTags{}
		if err := bus.Dispatch(c, getAllTags); err != nil {
			return c.Failure(err)
		}

		return c.Page(web.Props{
			Title:     "Manage Tags · Site Settings",
			ChunkName: "ManageTags.page",
			Data: web.Map{
				"tags": getAllTags.Result,
			},
		})
	}
}
