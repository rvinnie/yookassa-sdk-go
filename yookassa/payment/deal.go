// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

// The Deal within which the payment is being carried out.
type Deal struct {
	// Deal ID.
	ID string `json:"id,omitempty" binding:"min=36,max=50"`

	// Information about money distribution.
	Settlements []Settlement `json:"settlements,omitempty"`
}
