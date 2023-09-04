// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

type Recipient struct {
	// Store's ID in YooMoney.
	AccountID string `json:"account_id,omitempty"`

	// Subaccount's ID. Used for separating payment flows within one account.
	GatewayID string `json:"gateway_id,omitempty"`
}
