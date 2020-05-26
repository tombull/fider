package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/enum"

	"github.com/tombull/teamdream/app/models/query"

	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/dbx"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/errors"
	"github.com/tombull/teamdream/app/pkg/log"
	"github.com/tombull/teamdream/app/pkg/web"
)

//Health always returns OK
func Health() web.HandlerFunc {
	return func(c *web.Context) error {
		err := dbx.Ping()
		if err != nil {
			return c.Failure(err)
		}
		return c.Ok(web.Map{})
	}
}

//LegalPage returns a legal page with content from a file
func LegalPage(title, file string) web.HandlerFunc {
	return func(c *web.Context) error {
		bytes, err := ioutil.ReadFile(env.Etc(file))
		if err != nil {
			return c.NotFound()
		}

		return c.Render(http.StatusOK, "legal.html", web.Props{
			Title: title,
			Data: web.Map{
				"Content": string(bytes),
			},
		})
	}
}

//Sitemap returns the sitemap.xml of current site
func Sitemap() web.HandlerFunc {
	return func(c *web.Context) error {
		if c.Tenant().IsPrivate {
			return c.NotFound()
		}

		allPosts := &query.GetAllPosts{}
		if err := bus.Dispatch(c, allPosts); err != nil {
			return c.Failure(err)
		}

		baseURL := c.BaseURL()
		text := strings.Builder{}
		text.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
		text.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		text.WriteString(fmt.Sprintf("<url> <loc>%s</loc> </url>", baseURL))
		for _, post := range allPosts.Result {
			text.WriteString(fmt.Sprintf("<url> <loc>%s/posts/%d/%s</loc> </url>", baseURL, post.Number, post.Slug))
		}
		text.WriteString(`</urlset>`)

		c.Response.Header().Del("Content-Security-Policy")
		return c.XML(http.StatusOK, text.String())
	}
}

//RobotsTXT return content of robots.txt file
func RobotsTXT() web.HandlerFunc {
	return func(c *web.Context) error {
		bytes, err := ioutil.ReadFile(env.Path("./robots.txt"))
		if err != nil {
			return c.NotFound()
		}
		sitemapURL := c.BaseURL() + "/sitemap.xml"
		content := fmt.Sprintf("%s\nSitemap: %s", bytes, sitemapURL)
		return c.String(http.StatusOK, content)
	}
}

//Page returns a page without properties
func Page(title, description, chunkName string) web.HandlerFunc {
	return func(c *web.Context) error {
		return c.Page(web.Props{
			Title:       title,
			Description: description,
			ChunkName:   chunkName,
		})
	}
}

//BrowserNotSupported returns an error page for browser that Fider dosn't support
func BrowserNotSupported() web.HandlerFunc {
	return func(c *web.Context) error {
		return c.Render(http.StatusOK, "browser-not-supported.html", web.Props{
			Title:       "Browser not supported",
			Description: "We don't support this version of your browser",
		})
	}
}

//NewLogError is the input model for UI errors
type NewLogError struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//LogError logs an error coming from the UI
func LogError() web.HandlerFunc {
	return func(c *web.Context) error {
		input := new(NewLogError)
		err := c.Bind(input)
		if err != nil {
			return c.Failure(err)
		}
		log.Debugf(c, input.Message, dto.Props{
			"Data": input.Data,
		})
		return c.Ok(web.Map{})
	}
}

func validateKey(kind enum.EmailVerificationKind, c *web.Context) (*models.EmailVerification, error) {
	key := c.QueryParam("k")

	//If key has been used, return NotFound
	findByKey := &query.GetVerificationByKey{Kind: kind, Key: key}
	err := bus.Dispatch(c, findByKey)
	if err != nil {
		if errors.Cause(err) == app.ErrNotFound {
			return nil, c.NotFound()
		}
		return nil, c.Failure(err)
	}

	//If key has been used, return Gone
	if findByKey.Result.VerifiedAt != nil {
		return nil, c.Gone()
	}

	//If key expired, return Gone
	if time.Now().After(findByKey.Result.ExpiresAt) {
		err = bus.Dispatch(c, &cmd.SetKeyAsVerified{Key: key})
		if err != nil {
			return nil, c.Failure(err)
		}
		return nil, c.Gone()
	}

	return findByKey.Result, nil
}

func between(n, min, max int) int {
	if n > max {
		return max
	} else if n < min {
		return min
	}
	return n
}
