package actions_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app"

	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/bus"

	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/models"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/rand"
)

func TestCreateEditTag_InvalidName(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetTagBySlug) error {
		q.Result = &models.Tag{Slug: "feature-request", Name: "Feature Request", Color: "000000"}
		return nil
	})

	for _, name := range []string{
		"",
		"Feature Request",
		rand.String(31),
	} {
		action := &actions.CreateEditTag{Model: &models.CreateEditTag{Name: name, Color: "FFFFFF"}}
		result := action.Validate(context.Background(), nil)
		ExpectFailed(result, "name")
	}
}

func TestCreateEditTag_InvalidColor(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetTagBySlug) error {
		return app.ErrNotFound
	})

	for _, color := range []string{
		"",
		"ABC",
		"PPPOOO",
		"FFF",
		"000000X",
	} {
		action := &actions.CreateEditTag{Model: &models.CreateEditTag{Name: "Bug", Color: color}}
		result := action.Validate(context.Background(), nil)
		ExpectFailed(result, "color")
	}
}

func TestCreateEditTag_ValidInput(t *testing.T) {
	RegisterT(t)

	tag := &models.Tag{Slug: "to-discuss", Name: "To Discuss", Color: "000000"}
	bus.AddHandler(func(ctx context.Context, q *query.GetTagBySlug) error {
		if q.Slug == tag.Slug {
			q.Result = tag
			return nil
		} else {
			q.Result = nil
			return app.ErrNotFound
		}
	})

	action := &actions.CreateEditTag{Model: &models.CreateEditTag{Name: "Bug", Color: "FF0000"}}
	result := action.Validate(context.Background(), nil)
	ExpectSuccess(result)
	Expect(action.Tag).IsNil()

	action = &actions.CreateEditTag{Model: &models.CreateEditTag{Name: "New Name", Slug: "to-discuss", Color: "FF0000"}}
	result = action.Validate(context.Background(), nil)
	ExpectSuccess(result)
	Expect(action.Tag).Equals(tag)
}
