package postgres_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app"

	"github.com/tombull/teamdream/app/models"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/services/sqlstore/postgres"
)

func TestToTSQuery(t *testing.T) {
	RegisterT(t)

	var testcases = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"123 hello", "123|hello"},
		{" hello  ", "hello"},
		{" hello$ world$ ", "hello|world"},
		{" yes, please ", "yes|please"},
		{" yes / please ", "yes|please"},
		{" hello 'world' ", "hello|world"},
		{"hello|world", "hello|world"},
		{"hello | world", "hello|world"},
		{"hello & world", "hello|world"},
	}

	for _, testcase := range testcases {
		output := postgres.ToTSQuery(testcase.input)
		Expect(output).Equals(testcase.expected)
	}
}

func withTenant(ctx context.Context, tenant *models.Tenant) context.Context {
	return context.WithValue(ctx, app.TenantCtxKey, tenant)
}

func withUser(ctx context.Context, user *models.User) context.Context {
	ctx = context.WithValue(ctx, app.TenantCtxKey, user.Tenant)
	ctx = context.WithValue(ctx, app.UserCtxKey, user)
	return ctx
}
