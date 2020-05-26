package cmd

import (
	"testing"

	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/web"
)

func TestGetMainEngine(t *testing.T) {
	RegisterT(t)

	r := routes(web.New(nil))
	Expect(r).IsNotNil()
}
