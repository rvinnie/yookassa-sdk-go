// Package yoosettings describes all the necessary entities for working with YooMoney Settings.
package yoosettings

import (
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payout"
)

// Settings object contains relevant information about the configuration of the store or gateway.
type Settings struct {
	// Store's or gateway's ID in YooMoney.
	AccountID string `json:"account_id,omitempty"`

	// Store's or gateway status.
	Status *SettingsStatus `json:"status,omitempty"`

	// This is the Demo store or gateway.
	Test bool `json:"test,omitempty"`

	// Sending receipts enabled in store's or gateway's settings.
	FiscalizationEnabled bool `json:"fiscalization_enabled,omitempty"`

	// List of payment methods available to a store or gateway.
	PaymentMethods []yoopayment.PaymentMethoder `json:"payment_methods,omitempty"`

	// Store's INN (TIN): 10 or 12 digits.
	ITN string `json:"itn,omitempty"`

	// List of payout methods available to the gateway.
	PayoutMethods *[]yoopayout.PayoutMethodType `json:"payout_methods,omitempty"`

	// The name of the gateway, which is displayed in the YuKassa personal account.
	// Present if you requested gateway settings.
	Name string `json:"name,omitempty"`

	// The balance of your gateway. Present if you requested gateway settings.
	PayoutBalance *yoopayout.PayoutBalance `json:"payout_balance,omitempty"`
}
