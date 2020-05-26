package middlewares

import (
	"io"
	"time"

	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/pkg/errors"
	"github.com/tombull/teamdream/app/pkg/log"
	"github.com/tombull/teamdream/app/pkg/web"
)

func CatchPanic() web.MiddlewareFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(c *web.Context) error {
			defer func() {
				if r := recover(); r != nil {
					_ = c.Failure(errors.Panicked(r))
					c.Rollback()
					log.Infof(c, "@{HttpMethod:magenta} @{URL:magenta} panicked in @{ElapsedMs:magenta}ms (rolled back)", dto.Props{
						"HttpMethod": c.Request.Method,
						"URL":        c.Request.URL.String(),
						"ElapsedMs":  time.Since(c.Request.StartTime).Nanoseconds() / int64(time.Millisecond),
					})

					if f, ok := c.Response.(io.Closer); ok {
						f.Close()
					}
				}
			}()

			return next(c)
		}
	}
}
