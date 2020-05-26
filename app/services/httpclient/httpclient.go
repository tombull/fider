package httpclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/env"
)

func init() {
	http.DefaultClient.Timeout = 30 * time.Second
	bus.Register(Service{})
}

type Service struct{}

func (s Service) Name() string {
	return "HTTP"
}

func (s Service) Category() string {
	return "httpclient"
}

func (s Service) Enabled() bool {
	return !env.IsTest()
}

func (s Service) Init() {
	bus.AddHandler(requestHandler)
}

func requestHandler(ctx context.Context, c *cmd.HTTPRequest) error {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	if c.BasicAuth != nil {
		req.SetBasicAuth(c.BasicAuth.User, c.BasicAuth.Password)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.ResponseBody = respBody
	c.ResponseStatusCode = res.StatusCode
	return nil
}
