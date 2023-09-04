// Package yoorefund describes all the necessary entities for working with YooMoney Refunds.
package yoorefund

import "github.com/rvinnie/yookassa-sdk-go/yookassa/common"

// Information about money held for refunds: the amount to be held and the stores getting the refunds.
type Source struct {
	// ID of the store in favor of which you're accepting the receipt.
	// Provided by YooMoney, displayed in the Sellers section of your Merchant Profile (shopId column).
	AccountID string `json:"account_id,omitempty"`

	// Refund amount.
	Amount *yoocommon.Amount `json:"amount,omitempty"`

	// Commission that you charged during payment.
	PlatformFeeAmount *yoocommon.Amount `json:"platform_fee_amount,omitempty"`
}
