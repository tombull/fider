package cmd

import "github.com/tombull/teamdream/app/models/dto"

type CancelBillingSubscription struct {
}

type CreateBillingSubscription struct {
	PlanID string
}

type CreateBillingCustomer struct {
}

type DeleteBillingCustomer struct {
}

type ClearPaymentInfo struct {
}

type UpdatePaymentInfo struct {
	Input *dto.CreateEditBillingPaymentInfo
}
