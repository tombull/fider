package middlewares_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/tombull/teamdream/app/middlewares"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/mock"
	"github.com/tombull/teamdream/app/pkg/web"
)

func TestCache(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	server.Use(middlewares.ClientCache(5 * time.Minute))
	handler := func(c *web.Context) error {
		return c.NoContent(http.StatusOK)
	}

	status, response := server.Execute(handler)

	Expect(status).Equals(http.StatusOK)
	Expect(response.Header().Get("Cache-Control")).Equals("public, max-age=300")
}
