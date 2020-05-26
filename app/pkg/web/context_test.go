package web_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/tombull/teamdream/app/models"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/web"
)

func newGetContext(rawurl string, headers map[string]string) *web.Context {
	u, _ := url.Parse(rawurl)
	e := web.New(nil)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", u.RequestURI(), nil)
	req.Host = u.Host

	if u.Scheme == "https" {
		req.TLS = &tls.ConnectionState{}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return web.NewContext(e, req, res, nil)
}

func newBodyContext(method string, params web.StringMap, body, contentType string) *web.Context {
	e := web.New(nil)
	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, "/some/resource", strings.NewReader(body))
	req.Host = "demo.test.teamdream.co.uk:3000"
	req.Header.Set("Content-Type", contentType)
	return web.NewContext(e, req, res, params)
}

func TestContextID(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", nil)

	Expect(ctx.ContextID()).IsNotEmpty()
	Expect(ctx.ContextID()).HasLen(32)
}

func TestBaseURL(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", nil)

	Expect(ctx.BaseURL()).Equals("http://demo.test.teamdream.co.uk:3000")
}

func TestBaseURL_HTTPS(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("https://demo.test.teamdream.co.uk:3000", nil)

	Expect(ctx.BaseURL()).Equals("https://demo.test.teamdream.co.uk:3000")
}

func TestBaseURL_HTTPS_Proxy(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", map[string]string{
		"X-Forwarded-Proto": "https",
	})

	Expect(ctx.BaseURL()).Equals("https://demo.test.teamdream.co.uk:3000")
}

func TestCurrentURL(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000/resource?id=23", nil)

	Expect(ctx.Request.URL.String()).Equals("http://demo.test.teamdream.co.uk:3000/resource?id=23")
}

func TestTenantURL(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://login.test.teamdream.co.uk:3000", nil)
	tenant := &models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
	}
	Expect(web.TenantBaseURL(ctx, tenant)).Equals("http://theavengers.test.teamdream.co.uk:3000")
}

func TestTenantURL_WithCNAME(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", nil)
	tenant := &models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
		CNAME:     "feedback.theavengers.com",
	}
	Expect(web.TenantBaseURL(ctx, tenant)).Equals("http://feedback.theavengers.com:3000")
}

func TestTenantURL_SingleHostMode(t *testing.T) {
	RegisterT(t)
	env.Config.HostMode = "single"

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", nil)
	tenant := &models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
	}
	Expect(web.TenantBaseURL(ctx, tenant)).Equals("http://demo.test.teamdream.co.uk:3000")
}

func TestGlobalAssetsURL_SingleHostMode(t *testing.T) {
	RegisterT(t)

	env.Config.HostMode = "single"
	ctx := newGetContext("http://feedback.theavengers.com:3000", nil)
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.js")).Equals("http://feedback.theavengers.com:3000/assets/main.js")
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.css")).Equals("http://feedback.theavengers.com:3000/assets/main.css")

	env.Config.CDN.Host = "assets-teamdream.co.uk"
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.js")).Equals("http://assets-teamdream.co.uk/assets/main.js")
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.css")).Equals("http://assets-teamdream.co.uk/assets/main.css")
}

func TestGlobalAssetsURL_MultiHostMode(t *testing.T) {
	RegisterT(t)

	env.Config.HostMode = "multi"
	ctx := newGetContext("http://theavengers.test.teamdream.co.uk:3000", nil)
	ctx.SetTenant(&models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
		CNAME:     "feedback.theavengers.com",
	})

	Expect(web.GlobalAssetsURL(ctx, "/assets/main.js")).Equals("http://theavengers.test.teamdream.co.uk:3000/assets/main.js")
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.css")).Equals("http://theavengers.test.teamdream.co.uk:3000/assets/main.css")

	env.Config.CDN.Host = "assets-teamdream.co.uk"
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.js")).Equals("http://cdn.assets-teamdream.co.uk/assets/main.js")
	Expect(web.GlobalAssetsURL(ctx, "/assets/main.css")).Equals("http://cdn.assets-teamdream.co.uk/assets/main.css")
}

func TestTenantAssetsURL_SingleHostMode(t *testing.T) {
	RegisterT(t)

	env.Config.HostMode = "single"
	ctx := newGetContext("http://feedback.theavengers.com:3000", nil)
	ctx.SetTenant(&models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
	})

	Expect(web.TenantAssetsURL(ctx, "/assets/main.js")).Equals("http://feedback.theavengers.com:3000/assets/main.js")
	Expect(web.TenantAssetsURL(ctx, "/assets/main.css")).Equals("http://feedback.theavengers.com:3000/assets/main.css")

	env.Config.CDN.Host = "assets-teamdream.co.uk"
	Expect(web.TenantAssetsURL(ctx, "/assets/main.js")).Equals("http://assets-teamdream.co.uk/assets/main.js")
	Expect(web.TenantAssetsURL(ctx, "/assets/main.css")).Equals("http://assets-teamdream.co.uk/assets/main.css")
}

func TestTenantAssetsURL_MultiHostMode(t *testing.T) {
	RegisterT(t)

	env.Config.HostMode = "multi"
	ctx := newGetContext("http://theavengers.test.teamdream.co.uk:3000", nil)
	ctx.SetTenant(&models.Tenant{
		ID:        1,
		Subdomain: "theavengers",
		CNAME:     "feedback.theavengers.com",
	})

	Expect(web.TenantAssetsURL(ctx, "/assets/main.js")).Equals("http://theavengers.test.teamdream.co.uk:3000/assets/main.js")
	Expect(web.TenantAssetsURL(ctx, "/assets/main.css")).Equals("http://theavengers.test.teamdream.co.uk:3000/assets/main.css")

	env.Config.CDN.Host = "assets-teamdream.co.uk"
	Expect(web.TenantAssetsURL(ctx, "/assets/main.js")).Equals("http://theavengers.assets-teamdream.co.uk/assets/main.js")
	Expect(web.TenantAssetsURL(ctx, "/assets/main.css")).Equals("http://theavengers.assets-teamdream.co.uk/assets/main.css")
}

func TestCanonicalURL_SameDomain(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://theavengers.test.teamdream.co.uk:3000", nil)

	ctx.SetCanonicalURL("")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://theavengers.test.teamdream.co.uk:3000`)

	ctx.SetCanonicalURL("/some-url")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://theavengers.test.teamdream.co.uk:3000/some-url`)

	ctx.SetCanonicalURL("/some-other-url")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://theavengers.test.teamdream.co.uk:3000/some-other-url`)

	ctx.SetCanonicalURL("page-b/abc.html")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://theavengers.test.teamdream.co.uk:3000/page-b/abc.html`)
}

func TestCanonicalURL_DifferentDomain(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://theavengers.test.teamdream.co.uk:3000", nil)

	ctx.SetCanonicalURL("http://feedback.theavengers.com/some-url")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://feedback.theavengers.com/some-url`)

	ctx.SetCanonicalURL("")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://feedback.theavengers.com`)

	ctx.SetCanonicalURL("/some-other-url")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://feedback.theavengers.com/some-other-url`)

	ctx.SetCanonicalURL("page-b/abc.html")
	Expect(ctx.Value("Canonical-URL")).Equals(`http://feedback.theavengers.com/page-b/abc.html`)
}

func TestTryAgainLater(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000", nil)
	err := ctx.TryAgainLater(24 * time.Hour)
	Expect(err).IsNil()
	resp := ctx.Response.(*httptest.ResponseRecorder)
	Expect(ctx.ResponseStatusCode).Equals(http.StatusServiceUnavailable)
	Expect(resp.Code).Equals(http.StatusServiceUnavailable)
	Expect(resp.Header().Get("Cache-Control")).Equals("no-cache, no-store, must-revalidate")
	Expect(resp.Header().Get("Retry-After")).Equals("86400")
}

func TestGetOAuthBaseURL(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("https://mydomain.com/hello-world", nil)

	env.Config.HostMode = "multi"
	Expect(web.OAuthBaseURL(ctx)).Equals("https://login.test.teamdream.co.uk")

	env.Config.HostMode = "single"
	Expect(web.OAuthBaseURL(ctx)).Equals("https://mydomain.com")
}

func TestGetOAuthBaseURL_WithPort(t *testing.T) {
	RegisterT(t)

	ctx := newGetContext("http://demo.test.teamdream.co.uk:3000/hello-world", nil)

	env.Config.HostMode = "multi"
	Expect(web.OAuthBaseURL(ctx)).Equals("http://login.test.teamdream.co.uk:3000")

	env.Config.HostMode = "single"
	Expect(web.OAuthBaseURL(ctx)).Equals("http://demo.test.teamdream.co.uk:3000")
}
