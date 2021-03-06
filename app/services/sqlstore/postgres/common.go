package postgres

import (
	"context"
	"net/url"
	"regexp"
	"strings"

	"github.com/tombull/teamdream/app/models/enum"
	"github.com/tombull/teamdream/app/pkg/web"
)

var onlyalphanumeric = regexp.MustCompile("[^a-zA-Z0-9 |]+")
var replaceOr = strings.NewReplacer("|", " ")

// ToTSQuery converts input to another string that can be safely used for ts_query
func ToTSQuery(input string) string {
	input = replaceOr.Replace(onlyalphanumeric.ReplaceAllString(input, ""))
	return strings.Join(strings.Fields(input), "|")
}

func getViewData(view string) (string, []enum.PostStatus, string) {
	var (
		condition string
		sort      string
	)
	statuses := []enum.PostStatus{
		enum.PostOpen,
		enum.PostStarted,
		enum.PostPlanned,
	}
	switch view {
	case "recent":
		sort = "id"
	case "my-votes":
		condition = "AND has_voted = true"
		sort = "id"
	case "most-wanted":
		sort = "votes_count"
	case "most-discussed":
		sort = "comments_count"
	case "planned":
		sort = "response_date"
		statuses = []enum.PostStatus{enum.PostPlanned}
	case "started":
		sort = "response_date"
		statuses = []enum.PostStatus{enum.PostStarted}
	case "completed":
		sort = "response_date"
		statuses = []enum.PostStatus{enum.PostCompleted}
	case "declined":
		sort = "response_date"
		statuses = []enum.PostStatus{enum.PostDeclined}
	case "all":
		sort = "id"
		statuses = []enum.PostStatus{
			enum.PostOpen,
			enum.PostStarted,
			enum.PostPlanned,
			enum.PostCompleted,
			enum.PostDeclined,
		}
	case "trending":
		fallthrough
	default:
		sort = "CAST(((COALESCE(recent_votes_count, 0)*5 + COALESCE(recent_comments_count, 0) *3)-1) AS FLOAT) / pow((EXTRACT(EPOCH FROM current_timestamp - created_at)/3600) + 2, 1.4)"
	}
	return condition, statuses, sort
}

func buildAvatarURL(ctx context.Context, avatarType enum.AvatarType, id int, name, avatarBlobKey string) string {
	if name == "" {
		name = "-"
	}

	if avatarType == enum.AvatarTypeCustom {
		return web.TenantAssetsURL(ctx, "/images/%s", avatarBlobKey)
	} else {
		return web.TenantAssetsURL(ctx, "/avatars/%s/%d/%s", avatarType.String(), id, url.PathEscape(name))
	}
}
