// +build windows

package cmd

import (
	"os"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/pkg/web"
)

var extraSignals = []os.Signal{}

func handleExtraSignal(s os.Signal, e *web.Engine, settings *models.SystemSettings) int {
	return -1
}
