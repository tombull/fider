package main

import (
	"os"
	"runtime"

	_ "github.com/lib/pq"
	"github.com/tombull/teamdream/app/cmd"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/pkg/env"
)

var (
	// replaced during CI build
	buildtime   = ""
	buildnumber = "local"

	// Use this for non-stable releases
	// version = "x.y.z-" + buildnumber

	// Use this for stable releases
	// version = "x.y.z"

	version = "0.19.0-" + buildnumber
)

func main() {
	settings := &models.SystemSettings{
		BuildTime:       buildtime,
		Version:         version,
		Compiler:        runtime.Version(),
		Environment:     env.Config.Environment,
		GoogleAnalytics: env.Config.GoogleAnalytics,
		Mode:            env.Config.HostMode,
		Domain:          env.MultiTenantDomain(),
		HasLegal:        env.HasLegal(),
	}

	args := os.Args[1:]
	if len(args) > 0 && args[0] == "ping" {
		os.Exit(cmd.RunPing())
	} else if len(args) > 0 && args[0] == "migrate" {
		os.Exit(cmd.RunMigrate())
	} else {
		os.Exit(cmd.RunServer(settings))
	}
}
