package middlewares_test

import (
	"net/http"
	"testing"

	"github.com/tombull/teamdream/app/middlewares"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/mock"
	"github.com/tombull/teamdream/app/pkg/web"
)

func TestSecureWithoutCDN(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	server.Use(middlewares.Secure())

	var ctxID string
	status, response := server.Execute(func(c *web.Context) error {
		ctxID = c.ContextID()
		return c.NoContent(http.StatusOK)
	})

	expectedPolicy := "base-uri 'self'; default-src 'self'; style-src 'self' 'nonce-" + ctxID + "' https://fonts.googleapis.com ; script-src 'self' 'nonce-" + ctxID + "' https://cdn.polyfill.io https://js.stripe.com https://www.google-analytics.com ; img-src 'self' https: data: ; font-src 'self' https://fonts.gstatic.com data: ; object-src 'none'; media-src 'none'; connect-src 'self' https://www.google-analytics.com https://ipinfo.io https://js.stripe.com ; frame-src 'self' https://js.stripe.com"

	Expect(status).Equals(http.StatusOK)
	Expect(response.Header().Get("Content-Security-Policy")).Equals(expectedPolicy)
	Expect(response.Header().Get("X-XSS-Protection")).Equals("1; mode=block")
	Expect(response.Header().Get("X-Content-Type-Options")).Equals("nosniff")
	Expect(response.Header().Get("Referrer-Policy")).Equals("no-referrer-when-downgrade")
}

func TestSecureWithCDN(t *testing.T) {
	RegisterT(t)

	env.Config.CDN.Host = "test.fider.io"

	server := mock.NewServer()
	server.Use(middlewares.Secure())

	var ctxID string
	status, response := server.Execute(func(c *web.Context) error {
		ctxID = c.ContextID()
		return c.NoContent(http.StatusOK)
	})

	expectedPolicy := "base-uri 'self'; default-src 'self'; style-src 'self' 'nonce-" + ctxID + "' https://fonts.googleapis.com *.test.fider.io; script-src 'self' 'nonce-" + ctxID + "' https://cdn.polyfill.io https://js.stripe.com https://www.google-analytics.com *.test.fider.io; img-src 'self' https: data: *.test.fider.io; font-src 'self' https://fonts.gstatic.com data: *.test.fider.io; object-src 'none'; media-src 'none'; connect-src 'self' https://www.google-analytics.com https://ipinfo.io https://js.stripe.com *.test.fider.io; frame-src 'self' https://js.stripe.com"

	Expect(status).Equals(http.StatusOK)
	Expect(response.Header().Get("Content-Security-Policy")).Equals(expectedPolicy)
	Expect(response.Header().Get("X-XSS-Protection")).Equals("1; mode=block")
	Expect(response.Header().Get("X-Content-Type-Options")).Equals("nosniff")
	Expect(response.Header().Get("Referrer-Policy")).Equals("no-referrer-when-downgrade")
}

func TestSecureWithCDN_SingleHost(t *testing.T) {
	RegisterT(t)

	env.Config.CDN.Host = "test.fider.io"

	server := mock.NewSingleTenantServer()
	server.Use(middlewares.Secure())

	var ctxID string
	status, response := server.Execute(func(c *web.Context) error {
		ctxID = c.ContextID()
		return c.NoContent(http.StatusOK)
	})

	expectedPolicy := "base-uri 'self'; default-src 'self'; style-src 'self' 'nonce-" + ctxID + "' https://fonts.googleapis.com test.fider.io; script-src 'self' 'nonce-" + ctxID + "' https://cdn.polyfill.io https://js.stripe.com https://www.google-analytics.com test.fider.io; img-src 'self' https: data: test.fider.io; font-src 'self' https://fonts.gstatic.com data: test.fider.io; object-src 'none'; media-src 'none'; connect-src 'self' https://www.google-analytics.com https://ipinfo.io https://js.stripe.com test.fider.io; frame-src 'self' https://js.stripe.com"

	Expect(status).Equals(http.StatusOK)
	Expect(response.Header().Get("Content-Security-Policy")).Equals(expectedPolicy)
	Expect(response.Header().Get("X-XSS-Protection")).Equals("1; mode=block")
	Expect(response.Header().Get("X-Content-Type-Options")).Equals("nosniff")
	Expect(response.Header().Get("Referrer-Policy")).Equals("no-referrer-when-downgrade")
}
