package handlers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/bus"

	"github.com/tombull/teamdream/app/handlers"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/mock"
)

func TestHealthHandler(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	code, _ := server.
		Execute(handlers.Health())

	Expect(code).Equals(http.StatusOK)
}

func TestPageHandler(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	code, _ := server.
		Execute(handlers.Page("The Title", "The Description", "TheChunk.Page"))

	Expect(code).Equals(http.StatusOK)
}

func TestLegalPageHandler(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	code, _ := server.
		Execute(handlers.LegalPage("Terms of Service", "terms.md"))

	Expect(code).Equals(http.StatusOK)
}

func TestLegalPageHandler_Invalid(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	code, _ := server.
		Execute(handlers.LegalPage("Some Page", "somepage.md"))

	Expect(code).Equals(http.StatusNotFound)
}

func TestRobotsTXT(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	code, response := server.
		WithURL("https://demo.test.fider.io/robots.txt").
		Execute(handlers.RobotsTXT())
	content, _ := ioutil.ReadAll(response.Body)
	Expect(code).Equals(http.StatusOK)
	Expect(string(content)).Equals(`User-agent: *
Disallow: /_api/
Disallow: /api/v1/
Disallow: /admin/
Disallow: /oauth/
Sitemap: https://demo.test.fider.io/sitemap.xml`)
}

func TestSitemap(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetAllPosts) error {
		q.Result = []*models.Post{}
		return nil
	})

	server := mock.NewServer()
	code, response := server.
		OnTenant(mock.DemoTenant).
		WithURL("http://demo.test.fider.io:3000/sitemap.xml").
		Execute(handlers.Sitemap())

	bytes, _ := ioutil.ReadAll(response.Body)
	Expect(code).Equals(http.StatusOK)
	Expect(string(bytes)).Equals(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url> <loc>http://demo.test.fider.io:3000</loc> </url></urlset>`)
}

func TestSitemap_WithPosts(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetAllPosts) error {
		q.Result = []*models.Post{
			&models.Post{Number: 1, Slug: "my-new-idea-1", Title: "My new idea 1"},
			&models.Post{Number: 2, Slug: "the-other-idea", Title: "The other idea"},
		}
		return nil
	})

	server := mock.NewServer()

	code, response := server.
		OnTenant(mock.DemoTenant).
		WithURL("http://demo.test.fider.io:3000/sitemap.xml").
		Execute(handlers.Sitemap())

	bytes, _ := ioutil.ReadAll(response.Body)
	Expect(code).Equals(http.StatusOK)
	Expect(string(bytes)).Equals(`<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url> <loc>http://demo.test.fider.io:3000</loc> </url><url> <loc>http://demo.test.fider.io:3000/posts/1/my-new-idea-1</loc> </url><url> <loc>http://demo.test.fider.io:3000/posts/2/the-other-idea</loc> </url></urlset>`)
}

func TestSitemap_PrivateTenant_WithPosts(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()

	mock.DemoTenant.IsPrivate = true

	code, _ := server.
		OnTenant(mock.DemoTenant).
		WithURL("http://demo.test.fider.io:3000/sitemap.xml").
		Execute(handlers.Sitemap())

	Expect(code).Equals(http.StatusNotFound)
}
