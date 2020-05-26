package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/log"
	"github.com/tombull/teamdream/app/pkg/web"

	_ "github.com/tombull/teamdream/app/services/billing"
	_ "github.com/tombull/teamdream/app/services/blob/fs"
	_ "github.com/tombull/teamdream/app/services/blob/s3"
	_ "github.com/tombull/teamdream/app/services/blob/sql"
	_ "github.com/tombull/teamdream/app/services/email/mailgun"
	_ "github.com/tombull/teamdream/app/services/email/smtp"
	_ "github.com/tombull/teamdream/app/services/httpclient"
	_ "github.com/tombull/teamdream/app/services/log/console"
	_ "github.com/tombull/teamdream/app/services/log/sql"
	_ "github.com/tombull/teamdream/app/services/oauth"
	_ "github.com/tombull/teamdream/app/services/sqlstore/postgres"
)

//RunServer starts the Teamdream Server
//Returns an exitcode, 0 for OK and 1 for ERROR
func RunServer(settings *models.SystemSettings) int {
	svcs := bus.Init()
	ctx := log.WithProperty(context.Background(), log.PropertyKeyTag, "BOOTSTRAP")
	for _, s := range svcs {
		log.Debugf(ctx, "Service '@{ServiceCategory}.@{ServiceName}' has been initialized.", dto.Props{
			"ServiceCategory": s.Category(),
			"ServiceName":     s.Name(),
		})
	}

	bus.Publish(ctx, &cmd.PurgeExpiredNotifications{})

	e := routes(web.New(settings))

	go e.Start(":" + env.Config.Port)
	return listenSignals(e, settings)
}

func listenSignals(e *web.Engine, settings *models.SystemSettings) int {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, append([]os.Signal{syscall.SIGTERM, syscall.SIGINT}, extraSignals...)...)
	for {
		s := <-signals
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			err := e.Stop()
			if err != nil {
				return 1
			}
			return 0
		default:
			ret := handleExtraSignal(s, e, settings)
			if ret >= 0 {
				return ret
			}
		}
	}
}
