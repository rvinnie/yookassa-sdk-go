// Package yoopayout describes all the necessary entities for working with YooMoney Payouts.
package yoopayout

// The balance of your gateway.
type PayoutBalance struct {
	// Amount in the selected currency.
	// Expressed as a string and written with a period, for example 10.00.
	// The number of characters after the dot depends on the selected currency.
	Value string `json:"value,omitempty"`

	// Three-letter currency code in ISO-4217 format.
	// Example: RUB. Must match the subaccount currency (recipient.gateway_id) if you separate payment flows,
	// and the account currency (shopId in your account) if you don't.
	Currency string `json:"currency,omitempty"`
}
