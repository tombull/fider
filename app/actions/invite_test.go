package actions_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/query"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/bus"
)

func TestInviteUsers_Empty(t *testing.T) {
	RegisterT(t)

	action := &actions.InviteUsers{Model: &models.InviteUsers{}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "subject", "message", "recipients")
	Expect(action.Invitations).IsNil()
}

func TestInviteUsers_Oversized(t *testing.T) {
	RegisterT(t)

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject:    "Join us and share your ideas. Because we have a cool website and this subject needs to be very long",
		Message:    "Use this link to join %invite%",
		Recipients: []string{"jon.snow@got.com"},
	}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "subject")
	Expect(action.Invitations).IsNil()
}

func TestInviteUsers_MissingInvite(t *testing.T) {
	RegisterT(t)

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject:    "Share your feedback.",
		Message:    "Please!",
		Recipients: []string{"jon.snow@got.com"},
	}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "message")
	Expect(action.Invitations).IsNil()
}

func TestInviteUsers_TooManyRecipients(t *testing.T) {
	RegisterT(t)

	recipients := make([]string, 31)
	for i := 0; i < len(recipients); i++ {
		recipients[i] = fmt.Sprintf("jon.snow%d@got.com", i)
	}

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject:    "Share your feedback.",
		Message:    "Use this link to join %invite%",
		Recipients: recipients,
	}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "recipients")
	Expect(action.Invitations).IsNil()
}

func TestInviteUsers_InvalidRecipient(t *testing.T) {
	RegisterT(t)

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject: "Share your feedback.",
		Message: "Use this link to join our community: %invite%",
		Recipients: []string{
			"jon.snow",
			"@got.com",
		},
	}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "recipients")
	Expect(action.Invitations).IsNil()
}

func TestInviteUsers_Valid(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetUserByEmail) error {
		return app.ErrNotFound
	})

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject: "Share your feedback.",
		Message: "Use this link to join our community: %invite%",
		Recipients: []string{
			"",
			"jon.snow@got.com",
			"arya.stark@got.com",
		},
	}}

	ExpectSuccess(action.Validate(context.Background(), nil))

	Expect(action.Invitations).HasLen(2)

	Expect(action.Invitations[0].Email).Equals("jon.snow@got.com")
	Expect(action.Invitations[0].VerificationKey).IsNotEmpty()

	Expect(action.Invitations[1].Email).Equals("arya.stark@got.com")
	Expect(action.Invitations[1].VerificationKey).IsNotEmpty()
}

func TestInviteUsers_IgnoreAlreadyRegistered(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetUserByEmail) error {
		if q.Email == "tony.stark@avengers.com" {
			q.Result = &models.User{Email: q.Email}
			return nil
		}
		return app.ErrNotFound
	})

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject: "Share your feedback.",
		Message: "Use this link to join our community: %invite%",
		Recipients: []string{
			"tony.stark@avengers.com",
			"jon.snow@got.com",
			"arya.stark@got.com",
		},
	}}

	ExpectSuccess(action.Validate(context.Background(), nil))

	Expect(action.Invitations).HasLen(2)

	Expect(action.Invitations[0].Email).Equals("jon.snow@got.com")
	Expect(action.Invitations[0].VerificationKey).IsNotEmpty()

	Expect(action.Invitations[1].Email).Equals("arya.stark@got.com")
	Expect(action.Invitations[1].VerificationKey).IsNotEmpty()
}

func TestInviteUsers_ShouldFail_WhenAllRecipientsIgnored(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetUserByEmail) error {
		q.Result = &models.User{Email: q.Email}
		return nil
	})

	action := &actions.InviteUsers{Model: &models.InviteUsers{
		Subject: "Share your feedback.",
		Message: "Use this link to join our community: %invite%",
		Recipients: []string{
			"tony.stark@avengers.com",
		},
	}}

	ExpectFailed(action.Validate(context.Background(), nil), "recipients")
}

func TestInviteUsers_SampleInvite_IgnoreRecipients(t *testing.T) {
	RegisterT(t)

	action := &actions.InviteUsers{
		IsSampleInvite: true,
		Model: &models.InviteUsers{
			Subject: "Share your feedback.",
			Message: "Use this link to join our community: %invite%",
		},
	}

	ExpectSuccess(action.Validate(context.Background(), nil))
}
