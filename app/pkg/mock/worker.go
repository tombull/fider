package mock

import (
	"context"
	"net/url"

	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/pkg/web"
	"github.com/tombull/teamdream/app/pkg/worker"
)

// Worker is fake wrapper for background worker
type Worker struct {
	tenant  *models.Tenant
	user    *models.User
	baseURL string
}

func createWorker() *Worker {
	return &Worker{}
}

// OnTenant set current context tenant
func (w *Worker) OnTenant(tenant *models.Tenant) *Worker {
	w.tenant = tenant
	return w
}

// AsUser set current context user
func (w *Worker) AsUser(user *models.User) *Worker {
	w.user = user
	return w
}

// WithBaseURL set current context baseURL
func (w *Worker) WithBaseURL(baseURL string) *Worker {
	w.baseURL = baseURL
	return w
}

// Execute given task with current context
func (w *Worker) Execute(task worker.Task) error {
	task.OriginContext = context.Background()

	if w.user != nil {
		task.OriginContext = context.WithValue(task.OriginContext, app.UserCtxKey, w.user)
	}

	if w.tenant != nil {
		task.OriginContext = context.WithValue(task.OriginContext, app.TenantCtxKey, w.tenant)
	}

	u, _ := url.Parse(w.baseURL)
	if u != nil {
		task.OriginContext = context.WithValue(task.OriginContext, app.RequestCtxKey, web.Request{URL: u})
	}

	context := worker.NewContext(context.Background(), "0", task)
	return task.Job(context)
}

// NewNoopTask returns a worker task that does nothing
func NewNoopTask() worker.Task {
	return worker.Task{
		Name: "Noop Task",
		Job: func(c *worker.Context) error {
			return nil
		},
	}
}
