// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import (
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
)

// ReceiptInput contains information for Receipt creation.
type ReceiptInput struct {
	// Type of receipt in the online sales register: payment (payment) or payment refund (refund).
	Type ReceiptType `json:"type,omitempty"`

	// ID of the payment that the receipt was created for.
	PaymentID string `json:"payment_id,omitempty"`

	// ID of the refund that the receipt was created for. Not included in the payment receipt.
	RefundID string `json:"refund_id,omitempty"`

	// User details.
	Customer *Customer `json:"customer,omitempty"`

	// List of products in the receipt (maximum 100 items).
	Items []ReceiptItem `json:"items,omitempty"`

	// Creation of receipt in the online sales register immediately after receipt object creation.
	// You can only set true value at this time.
	Send bool `json:"send,omitempty"`

	// Store's taxation system.
	TaxSystemCode TaxSystemCode `json:"tax_system_code,omitempty"`

	// Additional user parameter (tag 1084 in 54-FZ).
	// You can specify this parameter if you're sending data
	// for creating a receipt using the Receipt after payment scenario.
	AdditionalUserProps *AdditionalUserProps `json:"additional_user_props,omitempty"`

	// Industry attribute of the payment subject (tag 1260 in 54-FZ).
	ReceiptIndustryDetails *ReceiptIndustryDetails `json:"receipt_industry_details,omitempty"`

	// Transaction attribute of the receipt (tag 1270 in 54-FZ).
	ReceiptOperationalDetails *ReceiptOperationalDetails `json:"receipt_operational_details,omitempty"`

	// List of completed payments.
	Settlements []yoocommon.Settlement `json:"settlements,omitempty"`

	// ID of the store on behalf of which you're sending the receipt.
	OnBehalfOf string `json:"on_behalf_of,omitempty"`
}
