package emailmock

import (
	"context"

	"github.com/tombull/teamdream/app"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/pkg/bus"
)

var MessageHistory = make([]*HistoryItem, 0)

type HistoryItem struct {
	From         string
	To           []dto.Recipient
	TemplateName string
	Props        dto.Props
	Tenant       *models.Tenant
}

func init() {
	bus.Register(Service{})
}

type Service struct{}

func (s Service) Name() string {
	return "Mock"
}

func (s Service) Category() string {
	return "email"
}

func (s Service) Enabled() bool {
	return true
}

func (s Service) Init() {
	MessageHistory = make([]*HistoryItem, 0)
	bus.AddListener(sendMail)
}

func sendMail(ctx context.Context, c *cmd.SendMail) {
	if c.Props == nil {
		c.Props = dto.Props{}
	}
	item := &HistoryItem{
		From:         c.From,
		To:           c.To,
		TemplateName: c.TemplateName,
		Props:        c.Props,
	}

	tenant, ok := ctx.Value(app.TenantCtxKey).(*models.Tenant)
	if ok {
		item.Tenant = tenant
	}
	MessageHistory = append(MessageHistory, item)
}
