package actions_test

import (
	"context"
	"testing"

	"github.com/tombull/teamdream/app/models/enum"
	"github.com/tombull/teamdream/app/models/query"

	"github.com/tombull/teamdream/app"
	"github.com/tombull/teamdream/app/actions"
	"github.com/tombull/teamdream/app/models"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/bus"
)

func TestSignInByEmail_EmptyEmail(t *testing.T) {
	RegisterT(t)

	action := actions.SignInByEmail{Model: &models.SignInByEmail{Email: " "}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "email")
}

func TestSignInByEmail_InvalidEmail(t *testing.T) {
	RegisterT(t)

	action := actions.SignInByEmail{Model: &models.SignInByEmail{Email: "Hi :)"}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "email")
}

func TestSignInByEmail_ShouldHaveVerificationKey(t *testing.T) {
	RegisterT(t)

	action := actions.SignInByEmail{}
	action.Initialize()
	action.Model.Email = "jon.snow@got.com"

	result := action.Validate(context.Background(), nil)
	ExpectSuccess(result)
	Expect(action.Model.VerificationKey).IsNotEmpty()
}

func TestCompleteProfile_EmptyNameAndKey(t *testing.T) {
	RegisterT(t)

	action := actions.CompleteProfile{Model: &models.CompleteProfile{}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "name", "key")
}

func TestCompleteProfile_LongName(t *testing.T) {
	RegisterT(t)

	action := actions.CompleteProfile{Model: &models.CompleteProfile{
		Name: "123456789012345678901234567890123456789012345678901", // 51 chars
	}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "name", "key")
}

func TestCompleteProfile_UnknownKey(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetVerificationByKey) error {
		return app.ErrNotFound
	})

	action := actions.CompleteProfile{Model: &models.CompleteProfile{Name: "Jon Snow", Key: "1234567890"}}
	result := action.Validate(context.Background(), nil)
	ExpectFailed(result, "key")
}

func TestCompleteProfile_ValidKey(t *testing.T) {
	RegisterT(t)

	key := "1234567890"
	bus.AddHandler(func(ctx context.Context, q *query.GetVerificationByKey) error {
		if q.Key == key && q.Kind == enum.EmailVerificationKindSignIn {
			q.Result = &models.EmailVerification{
				Key:   q.Key,
				Kind:  q.Kind,
				Email: "jon.snow@got.com",
			}
			return nil
		}
		return app.ErrNotFound
	})

	action := actions.CompleteProfile{Model: &models.CompleteProfile{Name: "Jon Snow", Key: key}}
	result := action.Validate(context.Background(), nil)

	ExpectSuccess(result)
	Expect(action.Model.Email).Equals("jon.snow@got.com")
}

func TestCompleteProfile_UserInvitation_ValidKey(t *testing.T) {
	RegisterT(t)

	key := "1234567890"
	bus.AddHandler(func(ctx context.Context, q *query.GetVerificationByKey) error {
		if q.Key == key && q.Kind == enum.EmailVerificationKindUserInvitation {
			q.Result = &models.EmailVerification{
				Key:   q.Key,
				Kind:  q.Kind,
				Email: "jon.snow@got.com",
			}
			return nil
		}
		return app.ErrNotFound
	})

	action := actions.CompleteProfile{Model: &models.CompleteProfile{Name: "Jon Snow", Key: key}}
	result := action.Validate(context.Background(), nil)

	ExpectSuccess(result)
	Expect(action.Model.Email).Equals("jon.snow@got.com")
}
