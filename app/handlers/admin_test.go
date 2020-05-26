package handlers_test

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/tombull/teamdream/app/models/cmd"

	"github.com/tombull/teamdream/app/models/query"
	. "github.com/tombull/teamdream/app/pkg/assert"
	"github.com/tombull/teamdream/app/pkg/bus"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/mock"
	"github.com/tombull/teamdream/app/services/blob/fs"

	"github.com/tombull/teamdream/app/handlers"
)

func TestUpdateSettingsHandler(t *testing.T) {
	RegisterT(t)

	var updateCmd *cmd.UpdateTenantSettings
	bus.AddHandler(func(ctx context.Context, c *cmd.UpdateTenantSettings) error {
		updateCmd = c
		return nil
	})

	bus.AddHandler(func(ctx context.Context, c *cmd.UploadImage) error {
		return nil
	})

	server := mock.NewServer()
	mock.DemoTenant.LogoBlobKey = "logos/hello-world.png"

	code, _ := server.
		OnTenant(mock.DemoTenant).
		AsUser(mock.JonSnow).
		ExecutePost(
			handlers.UpdateSettings(),
			`{ "title": "GoT", "invitation": "Join us!", "welcomeMessage": "Welcome to GoT Feedback Forum" }`,
		)

	Expect(code).Equals(http.StatusOK)
	Expect(updateCmd.Settings.Title).Equals("GoT")
	Expect(updateCmd.Settings.Invitation).Equals("Join us!")
	Expect(updateCmd.Settings.WelcomeMessage).Equals("Welcome to GoT Feedback Forum")
	Expect(updateCmd.Settings.Logo.BlobKey).Equals("logos/hello-world.png")
}

func TestUpdateSettingsHandler_NewLogo(t *testing.T) {
	RegisterT(t)
	bus.Init(fs.Service{})

	var updateCmd *cmd.UpdateTenantSettings
	bus.AddHandler(func(ctx context.Context, c *cmd.UpdateTenantSettings) error {
		updateCmd = c
		return nil
	})

	bus.AddHandler(func(ctx context.Context, c *cmd.UploadImage) error {
		c.Image.BlobKey = c.Folder + "/" + c.Image.Upload.FileName
		return nil
	})

	logoBytes, _ := ioutil.ReadFile(env.Etc("logo.png"))
	logoB64 := base64.StdEncoding.EncodeToString(logoBytes)

	server := mock.NewServer()
	code, _ := server.
		OnTenant(mock.DemoTenant).
		AsUser(mock.JonSnow).
		ExecutePost(
			handlers.UpdateSettings(), `{
				"title": "GoT",
				"invitation": "Join us!",
				"welcomeMessage": "Welcome to GoT Feedback Forum",
				"logo": {
					"upload": {
						"fileName": "picture.png",
						"contentType": "image/png",
						"content": "`+logoB64+`"
					}
				}
			}`)

	Expect(code).Equals(http.StatusOK)
	Expect(updateCmd.Settings.Title).Equals("GoT")
	Expect(updateCmd.Settings.Invitation).Equals("Join us!")
	Expect(updateCmd.Settings.WelcomeMessage).Equals("Welcome to GoT Feedback Forum")
	Expect(updateCmd.Settings.Logo.BlobKey).Equals("logos/picture.png")
}

func TestUpdateSettingsHandler_RemoveLogo(t *testing.T) {
	RegisterT(t)

	var updateCmd *cmd.UpdateTenantSettings
	bus.AddHandler(func(ctx context.Context, c *cmd.UpdateTenantSettings) error {
		updateCmd = c
		return nil
	})

	bus.AddHandler(func(ctx context.Context, c *cmd.UploadImage) error {
		return nil
	})

	server := mock.NewServer()
	mock.DemoTenant.LogoBlobKey = "logos/hello-world.png"

	code, _ := server.
		OnTenant(mock.DemoTenant).
		AsUser(mock.JonSnow).
		ExecutePost(
			handlers.UpdateSettings(), `{
				"title": "GoT",
				"invitation": "Join us!",
				"welcomeMessage": "Welcome to GoT Feedback Forum",
				"logo": {
					"remove": true
				}
			}`)

	Expect(code).Equals(http.StatusOK)
	Expect(updateCmd.Settings.Title).Equals("GoT")
	Expect(updateCmd.Settings.Invitation).Equals("Join us!")
	Expect(updateCmd.Settings.WelcomeMessage).Equals("Welcome to GoT Feedback Forum")
	Expect(updateCmd.Settings.Logo.Remove).IsTrue()
}

func TestUpdatePrivacyHandler(t *testing.T) {
	RegisterT(t)

	var updateCmd *cmd.UpdateTenantPrivacySettings
	bus.AddHandler(func(ctx context.Context, c *cmd.UpdateTenantPrivacySettings) error {
		updateCmd = c
		return nil
	})

	server := mock.NewServer()
	code, _ := server.
		OnTenant(mock.DemoTenant).
		AsUser(mock.JonSnow).
		ExecutePost(
			handlers.UpdatePrivacy(),
			`{ "isPrivate": true }`,
		)

	Expect(code).Equals(http.StatusOK)
	Expect(updateCmd.Settings.IsPrivate).IsTrue()
}

func TestManageMembersHandler(t *testing.T) {
	RegisterT(t)

	bus.AddHandler(func(ctx context.Context, q *query.GetAllUsers) error {
		return nil
	})

	server := mock.NewServer()
	code, _ := server.
		OnTenant(mock.DemoTenant).
		Execute(
			handlers.ManageMembers(),
		)

	Expect(code).Equals(http.StatusOK)
}
