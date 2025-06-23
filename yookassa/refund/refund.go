// Package yoorefund describes all the necessary entities for working with YooMoney Refunds.
package yoorefund

import (
	"time"

	"github.com/imgrigorev/yookassa-sdk-go/yookassa/common"
)

// The Refund object contains all currently relevant information about the refund of a successful payment.
// The object is sent in response to any refund-related requests.
type Refund struct {
	// Refund's ID in YooMoney.
	Id string `json:"id,omitempty"`

	// Payment ID in YooMoney.
	PaymentId string `json:"payment_id,omitempty"`

	// Refund status. Possible values: pending, succeeded, and canceled.
	Status Status `json:"status,omitempty"`

	// Comment to the canceled status: who canceled the refund and what was the reason.
	CancellationDetails *yoocommon.CancellationDetails `json:"cancellation_details,omitempty"`

	// Status of receipt delivery. Possible values: pending, succeeded, and canceled.
	ReceiptRegistration Status `json:"receipt_registration,omitempty"`

	// Time to refund creation, based on UTC and specified in the ISO 8601 format,
	// for example, 2017-11-03T11:52:31.827Z
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Amount refunded to the user.
	Amount *yoocommon.Amount `json:"amount,omitempty"`

	// Reason behind the refund.
	Description string `json:"description,omitempty"`

	// Information about money held for refunds:
	// the amount to be held and the stores getting the refunds.
	Sources *Source `json:"sources,omitempty"`

	// The deal within which the refund is being carried out.
	Deal *Deal `json:"deal,omitempty"`
}
