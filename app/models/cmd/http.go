package cmd

import (
	"io"

	"github.com/tombull/teamdream/app/models/dto"
)

type HTTPRequest struct {
	URL       string
	Body      io.Reader
	Method    string
	Headers   map[string]string
	BasicAuth *dto.BasicAuth

	ResponseBody       []byte
	ResponseStatusCode int
}
