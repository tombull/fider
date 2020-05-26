package middlewares

import (
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/web"
)

// RequireBillingEnabled returns 404 if billing is not enabled, otherwise it continues the chain
func RequireBillingEnabled() web.MiddlewareFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(c *web.Context) error {
			if !env.IsBillingEnabled() || c.Tenant().Billing == nil {
				return c.NotFound()
			}
			return next(c)
		}
	}
}
