package middlewares

import (
	"strings"
	"time"

	"github.com/tombull/teamdream/app/pkg/rand"
	"github.com/tombull/teamdream/app/pkg/web"
)

// Session starts a new Session if an Session ID is not yet set
func Session() web.MiddlewareFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(c *web.Context) error {
			cookie, err := c.Request.Cookie(web.CookieSessionName)
			if err != nil {
				nextYear := time.Now().Add(365 * 24 * time.Hour)
				cookie = c.AddCookie(web.CookieSessionName, rand.String(48), nextYear)
			}
			c.SetSessionID(cookie.Value)
			err = next(c)
			cc := c.Response.Header().Get("Cache-Control")
			if strings.Contains(cc, "max-age=") {
				c.Response.Header().Del("Set-Cookie")
			}
			return err
		}
	}
}
