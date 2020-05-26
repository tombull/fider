package postgres

import (
	"context"
	"time"

	"github.com/tombull/teamdream/app/pkg/bus"

	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/enum"
	"github.com/tombull/teamdream/app/models/query"

	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/pkg/dbx"
	"github.com/tombull/teamdream/app/pkg/env"
	"github.com/tombull/teamdream/app/pkg/errors"
)

type dbTenant struct {
	ID             int              `db:"id"`
	Name           string           `db:"name"`
	Subdomain      string           `db:"subdomain"`
	CNAME          string           `db:"cname"`
	Invitation     string           `db:"invitation"`
	WelcomeMessage string           `db:"welcome_message"`
	Status         int              `db:"status"`
	IsPrivate      bool             `db:"is_private"`
	LogoBlobKey    string           `db:"logo_bkey"`
	CustomCSS      string           `db:"custom_css"`
	Billing        *dbTenantBilling `db:"billing"`
}

func (t *dbTenant) toModel() *models.Tenant {
	if t == nil {
		return nil
	}

	tenant := &models.Tenant{
		ID:             t.ID,
		Name:           t.Name,
		Subdomain:      t.Subdomain,
		CNAME:          t.CNAME,
		Invitation:     t.Invitation,
		WelcomeMessage: t.WelcomeMessage,
		Status:         t.Status,
		IsPrivate:      t.IsPrivate,
		LogoBlobKey:    t.LogoBlobKey,
		CustomCSS:      t.CustomCSS,
	}

	if t.Billing != nil && t.Billing.TrialEndsAt.Valid {
		tenant.Billing = &models.TenantBilling{
			TrialEndsAt:          t.Billing.TrialEndsAt.Time,
			StripeCustomerID:     t.Billing.StripeCustomerID.String,
			StripeSubscriptionID: t.Billing.StripeSubscriptionID.String,
			StripePlanID:         t.Billing.StripePlanID.String,
		}
		if t.Billing.SubscriptionEndsAt.Valid {
			tenant.Billing.SubscriptionEndsAt = &t.Billing.SubscriptionEndsAt.Time
		}
	}

	return tenant
}

type dbTenantBilling struct {
	StripeCustomerID     dbx.NullString `db:"stripe_customer_id"`
	StripeSubscriptionID dbx.NullString `db:"stripe_subscription_id"`
	StripePlanID         dbx.NullString `db:"stripe_plan_id"`
	TrialEndsAt          dbx.NullTime   `db:"trial_ends_at"`
	SubscriptionEndsAt   dbx.NullTime   `db:"subscription_ends_at"`
}

type dbEmailVerification struct {
	ID         int                        `db:"id"`
	Name       string                     `db:"name"`
	Email      string                     `db:"email"`
	Key        string                     `db:"key"`
	Kind       enum.EmailVerificationKind `db:"kind"`
	UserID     dbx.NullInt                `db:"user_id"`
	CreatedAt  time.Time                  `db:"created_at"`
	ExpiresAt  time.Time                  `db:"expires_at"`
	VerifiedAt dbx.NullTime               `db:"verified_at"`
}

func (t *dbEmailVerification) toModel() *models.EmailVerification {
	model := &models.EmailVerification{
		Name:       t.Name,
		Email:      t.Email,
		Key:        t.Key,
		Kind:       t.Kind,
		CreatedAt:  t.CreatedAt,
		ExpiresAt:  t.ExpiresAt,
		VerifiedAt: nil,
	}

	if t.VerifiedAt.Valid {
		model.VerifiedAt = &t.VerifiedAt.Time
	}

	if t.UserID.Valid {
		model.UserID = int(t.UserID.Int64)
	}

	return model
}

func isCNAMEAvailable(ctx context.Context, q *query.IsCNAMEAvailable) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		exists, err := trx.Exists("SELECT id FROM tenants WHERE cname = $1 AND id <> $2", q.CNAME, tenant.ID)
		if err != nil {
			q.Result = false
			return errors.Wrap(err, "failed to check if tenant exists with CNAME '%s'", q.CNAME)
		}
		q.Result = !exists
		return nil
	})
}

func isSubdomainAvailable(ctx context.Context, q *query.IsSubdomainAvailable) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		exists, err := trx.Exists("SELECT id FROM tenants WHERE subdomain = $1", q.Subdomain)
		if err != nil {
			q.Result = false
			return errors.Wrap(err, "failed to check if tenant exists with subdomain '%s'", q.Subdomain)
		}
		q.Result = !exists
		return nil
	})
}

func updateTenantPrivacySettings(ctx context.Context, c *cmd.UpdateTenantPrivacySettings) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		_, err := trx.Execute("UPDATE tenants SET is_private = $1 WHERE id = $2", c.Settings.IsPrivate, tenant.ID)
		if err != nil {
			return errors.Wrap(err, "failed update tenant privacy settings")
		}
		return nil
	})
}

func updateTenantSettings(ctx context.Context, c *cmd.UpdateTenantSettings) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		if c.Settings.Logo.Remove {
			c.Settings.Logo.BlobKey = ""
		}

		query := "UPDATE tenants SET name = $1, invitation = $2, welcome_message = $3, cname = $4, logo_bkey = $5 WHERE id = $6"
		_, err := trx.Execute(query, c.Settings.Title, c.Settings.Invitation, c.Settings.WelcomeMessage, c.Settings.CNAME, c.Settings.Logo.BlobKey, tenant.ID)
		if err != nil {
			return errors.Wrap(err, "failed update tenant settings")
		}

		tenant.Name = c.Settings.Title
		tenant.Invitation = c.Settings.Invitation
		tenant.CNAME = c.Settings.CNAME
		tenant.WelcomeMessage = c.Settings.WelcomeMessage

		return nil
	})
}

func updateTenantBillingSettings(ctx context.Context, c *cmd.UpdateTenantBillingSettings) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		_, err := trx.Execute(`
			UPDATE tenants_billing 
			SET stripe_customer_id = $1, 
					stripe_plan_id = $2, 
					stripe_subscription_id = $3, 
					subscription_ends_at = $4 
			WHERE tenant_id = $5
		`,
			c.Settings.StripeCustomerID,
			c.Settings.StripePlanID,
			c.Settings.StripeSubscriptionID,
			c.Settings.SubscriptionEndsAt,
			tenant.ID,
		)
		if err != nil {
			return errors.Wrap(err, "failed update tenant billing settings")
		}
		return nil
	})
}

func updateTenantAdvancedSettings(ctx context.Context, c *cmd.UpdateTenantAdvancedSettings) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		query := "UPDATE tenants SET custom_css = $1 WHERE id = $2"
		_, err := trx.Execute(query, c.Settings.CustomCSS, tenant.ID)
		if err != nil {
			return errors.Wrap(err, "failed update tenant advanced settings")
		}

		tenant.CustomCSS = c.Settings.CustomCSS
		return nil
	})
}

func activateTenant(ctx context.Context, c *cmd.ActivateTenant) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		query := "UPDATE tenants SET status = $1 WHERE id = $2"
		_, err := trx.Execute(query, enum.TenantActive, c.TenantID)
		if err != nil {
			return errors.Wrap(err, "failed to activate tenant with id '%d'", c.TenantID)
		}
		return nil
	})
}

func getVerificationByKey(ctx context.Context, q *query.GetVerificationByKey) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		verification := dbEmailVerification{}

		query := "SELECT id, email, name, key, created_at, verified_at, expires_at, kind, user_id FROM email_verifications WHERE key = $1 AND kind = $2 LIMIT 1"
		err := trx.Get(&verification, query, q.Key, q.Kind)
		if err != nil {
			return errors.Wrap(err, "failed to get email verification by its key")
		}

		q.Result = verification.toModel()
		return nil
	})
}

func saveVerificationKey(ctx context.Context, c *cmd.SaveVerificationKey) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		var userID interface{}
		if c.Request.GetUser() != nil {
			userID = c.Request.GetUser().ID
		}

		query := "INSERT INTO email_verifications (tenant_id, email, created_at, expires_at, key, name, kind, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
		_, err := trx.Execute(query, tenant.ID, c.Request.GetEmail(), time.Now(), time.Now().Add(c.Duration), c.Key, c.Request.GetName(), c.Request.GetKind(), userID)
		if err != nil {
			return errors.Wrap(err, "failed to save verification key for kind '%d'", c.Request.GetKind())
		}
		return nil
	})
}

func setKeyAsVerified(ctx context.Context, c *cmd.SetKeyAsVerified) error {
	return using(ctx, func(trx *dbx.Trx, tenant *models.Tenant, user *models.User) error {
		query := "UPDATE email_verifications SET verified_at = $1 WHERE tenant_id = $2 AND key = $3"
		_, err := trx.Execute(query, time.Now(), tenant.ID, c.Key)
		if err != nil {
			return errors.Wrap(err, "failed to update verified date of email verification request")
		}
		return nil
	})
}

func createTenant(ctx context.Context, c *cmd.CreateTenant) error {
	return using(ctx, func(trx *dbx.Trx, _ *models.Tenant, _ *models.User) error {
		now := time.Now()

		var id int
		err := trx.Get(&id,
			`INSERT INTO tenants (name, subdomain, created_at, cname, invitation, welcome_message, status, is_private, custom_css, logo_bkey) 
			 VALUES ($1, $2, $3, '', '', '', $4, false, '', '') 
			 RETURNING id`, c.Name, c.Subdomain, now, c.Status)
		if err != nil {
			return err
		}

		if env.IsBillingEnabled() {
			_, err = trx.Execute(
				`INSERT INTO tenants_billing (tenant_id, trial_ends_at) VALUES ($1, $2)`,
				id, now.Add(30*24*time.Hour),
			)
			if err != nil {
				return err
			}
		}

		byDomain := &query.GetTenantByDomain{Domain: c.Subdomain}
		err = bus.Dispatch(ctx, byDomain)
		c.Result = byDomain.Result
		return err
	})
}

func getFirstTenant(ctx context.Context, q *query.GetFirstTenant) error {
	return using(ctx, func(trx *dbx.Trx, _ *models.Tenant, _ *models.User) error {
		tenant := dbTenant{}

		err := trx.Get(&tenant, `
			SELECT t.id, t.name, t.subdomain, t.cname, t.invitation, t.welcome_message, t.status, t.is_private, t.logo_bkey, t.custom_css,
						 tb.trial_ends_at AS billing_trial_ends_at,
						 tb.subscription_ends_at AS billing_subscription_ends_at,
						 tb.stripe_customer_id AS billing_stripe_customer_id,
						 tb.stripe_plan_id AS billing_stripe_plan_id,
						 tb.stripe_subscription_id AS billing_stripe_subscription_id
			FROM tenants t
			LEFT JOIN tenants_billing tb
			ON tb.tenant_id = t.id
			ORDER BY t.id LIMIT 1
		`)

		if err != nil {
			return errors.Wrap(err, "failed to get first tenant")
		}

		q.Result = tenant.toModel()
		return nil
	})
}

func getTenantByDomain(ctx context.Context, q *query.GetTenantByDomain) error {
	return using(ctx, func(trx *dbx.Trx, _ *models.Tenant, _ *models.User) error {
		tenant := dbTenant{}

		err := trx.Get(&tenant, `
			SELECT t.id, t.name, t.subdomain, t.cname, t.invitation, t.welcome_message, t.status, t.is_private, t.logo_bkey, t.custom_css,
						 tb.trial_ends_at AS billing_trial_ends_at,
						 tb.subscription_ends_at AS billing_subscription_ends_at,
						 tb.stripe_customer_id AS billing_stripe_customer_id,
						 tb.stripe_plan_id AS billing_stripe_plan_id,
						 tb.stripe_subscription_id AS billing_stripe_subscription_id
			FROM tenants t
			LEFT JOIN tenants_billing tb
			ON tb.tenant_id = t.id
			WHERE t.subdomain = $1 OR t.subdomain = $2 OR t.cname = $3 
			ORDER BY t.cname DESC
		`, env.Subdomain(q.Domain), q.Domain, q.Domain)
		if err != nil {
			return errors.Wrap(err, "failed to get tenant with domain '%s'", q.Domain)
		}

		q.Result = tenant.toModel()
		return nil
	})
}
