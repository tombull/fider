package validate_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app/models/query"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/rand"
	"github.com/tombull/teamdream/app/pkg/validate"
)

func TestInvalidEmail(t *testing.T) {
	RegisterT(t)

	for _, email := range []string{
		"hello",
		"",
		"my@company",
		"my @company.com",
		"my@.company.com",
		"my+company.com",
		".my@company.com",
		"my@company@other.com",
		"my@company@other.com",
		rand.String(200) + "@gmail.com",
	} {
		messages := validate.Email(email)
		Expect(len(messages) > 0).IsTrue()
	}
}

func TestValidEmail(t *testing.T) {
	RegisterT(t)

	for _, email := range []string{
		"hello@company.com",
		"hello+alias@company.com",
		"abc@gmail.com",
	} {
		messages := validate.Email(email)
		Expect(messages).HasLen(0)
	}
}

func TestInvalidURL(t *testing.T) {
	RegisterT(t)

	for _, rawurl := range []string{
		"http//google.com",
		"google.com",
		"google",
		rand.String(301),
		"my@company",
	} {
		messages := validate.URL(rawurl)
		Expect(len(messages) > 0).IsTrue()
	}
}

func TestValidURL(t *testing.T) {
	RegisterT(t)

	for _, rawurl := range []string{
		"http://example.org",
		"https://example.org/oauth",
		"https://example.org/oauth?test=abc",
	} {
		messages := validate.URL(rawurl)
		Expect(messages).HasLen(0)
	}
}

func TestInvalidCNAME(t *testing.T) {
	RegisterT(t)

	for _, cname := range []string{
		"hello",
		"hellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohellohello.com",
		"",
		"my",
		"name.com/abc",
		"feedback.test.teamdream.co.uk",
		"test.teamdream.co.uk",
		"@google.com",
	} {
		messages := validate.CNAME(context.Background(), cname)
		Expect(len(messages) > 0).IsTrue()
	}
}

func TestValidHostname(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.IsCNAMEAvailable) error {
		q.Result = true
		return nil
	})

	for _, cname := range []string{
		"google.com",
		"feedback.teamdream.co.uk",
		"my.super.domain.com",
		"jon-snow.got.com",
		"got.com",
		"hi.m",
	} {
		messages := validate.CNAME(context.Background(), cname)
		Expect(messages).HasLen(0)
	}
}

func TestValidCNAME_Availability(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.IsCNAMEAvailable) error {
		q.Result = q.CNAME != "footbook.com" && q.CNAME != "teamdream.yourcompany.com" && q.CNAME != "feedback.newyork.com"
		return nil
	})

	for _, cname := range []string{
		"footbook.com",
		"teamdream.yourcompany.com",
		"feedback.newyork.com",
	} {
		messages := validate.CNAME(context.Background(), cname)
		Expect(len(messages) > 0).IsTrue()
	}

	for _, cname := range []string{
		"teamdream.footbook.com",
		"yourcompany.com",
		"anything.com",
	} {
		messages := validate.CNAME(context.Background(), cname)
		Expect(messages).HasLen(0)
	}
}
