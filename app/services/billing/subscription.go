package billing

import (
	"context"
	"strings"
	"time"

	"github.com/stripe/stripe-go"
	"github.com/tombull/teamdream/app/models"
	"github.com/tombull/teamdream/app/models/cmd"
	"github.com/tombull/teamdream/app/models/dto"
	"github.com/tombull/teamdream/app/models/query"
	"github.com/tombull/teamdream/app/pkg/errors"
)

func cancelSubscription(ctx context.Context, c *cmd.CancelBillingSubscription) error {
	return using(ctx, func(tenant *models.Tenant) error {
		sub, err := stripeClient.Subscriptions.Update(tenant.Billing.StripeSubscriptionID, &stripe.SubscriptionParams{
			CancelAtPeriodEnd: stripe.Bool(true),
		})
		if err != nil {
			return errors.Wrap(err, "failed to cancel stripe subscription")
		}
		endDate := time.Unix(sub.CurrentPeriodEnd, 0)
		tenant.Billing.SubscriptionEndsAt = &endDate
		return nil
	})
}

func subscribe(ctx context.Context, c *cmd.CreateBillingSubscription) error {
	return using(ctx, func(tenant *models.Tenant) error {
		customerID := tenant.Billing.StripeCustomerID
		if tenant.Billing.StripeSubscriptionID != "" {
			sub, err := stripeClient.Subscriptions.Get(tenant.Billing.StripeSubscriptionID, nil)
			if err != nil {
				return errors.Wrap(err, "failed to get stripe subscription")
			}
			_, err = stripeClient.Subscriptions.Update(tenant.Billing.StripeSubscriptionID, &stripe.SubscriptionParams{
				CancelAtPeriodEnd: stripe.Bool(false),
				Items: []*stripe.SubscriptionItemsParams{
					{
						ID:   stripe.String(sub.Items.Data[0].ID),
						Plan: stripe.String(c.PlanID),
					},
				},
			})

			if err != nil {
				return errors.Wrap(err, "failed to update stripe subscription")
			}

			tenant.Billing.SubscriptionEndsAt = nil
		} else {
			sub, err := stripeClient.Subscriptions.New(&stripe.SubscriptionParams{
				Customer: stripe.String(customerID),
				Items: []*stripe.SubscriptionItemsParams{
					{
						Plan: stripe.String(c.PlanID),
					},
				},
			})

			if err != nil {
				return errors.Wrap(err, "failed to create stripe subscription")
			}

			tenant.Billing.StripeSubscriptionID = sub.ID
		}

		tenant.Billing.StripePlanID = c.PlanID
		return nil
	})
}

func getUpcomingInvoice(ctx context.Context, q *query.GetUpcomingInvoice) error {
	return using(ctx, func(tenant *models.Tenant) error {
		inv, err := stripeClient.Invoices.GetNext(&stripe.InvoiceParams{
			Customer:     stripe.String(tenant.Billing.StripeCustomerID),
			Subscription: stripe.String(tenant.Billing.StripeSubscriptionID),
		})
		if err != nil {
			if stripeErr, ok := err.(*stripe.Error); ok {
				if stripeErr.HTTPStatusCode == 404 {
					return nil
				}
			}
			return errors.Wrap(err, "failed to get upcoming invoice")
		}

		dueDate := time.Unix(inv.DueDate, 0)
		q.Result = &dto.UpcomingInvoice{
			Currency:  strings.ToUpper(string(inv.Currency)),
			AmountDue: inv.AmountDue,
			DueDate:   dueDate,
		}
		return nil
	})
}
