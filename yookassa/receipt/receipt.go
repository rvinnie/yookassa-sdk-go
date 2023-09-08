// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import (
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	"time"
)

// Receipt contains up-to-date information about the receipt created for payment or refund.
type Receipt struct {
	// Receipt's ID in YooMoney.
	ID string `json:"id,omitempty"`

	// Type of receipt in the online sales register: payment (payment) or payment refund (refund).
	Type ReceiptType `json:"type,omitempty"`

	// ID of the payment that the receipt was created for.
	PaymentID string `json:"payment_id,omitempty"`

	// ID of the refund that the receipt was created for. Not included in the payment receipt.
	RefundID string `json:"refund_id,omitempty"`

	// Status of receipt delivery.
	Status ReceiptStatus `json:"status,omitempty"`

	// Fiscal document number.
	FiscalDocumentNumber string `json:"fiscal_document_number,omitempty"`

	// Number of fiscal storage drive in online sales register.
	FiscalStorageNumber string `json:"fiscal_storage_number,omitempty"`

	// Fiscal attribute of the receipt.
	// Created by the fiscal storage drive from the data sent for receipt registration.
	FiscalAttribute string `json:"fiscal_attribute,omitempty"`

	// Date and time of receipt creation in the fiscal storage drive, specified in the ISO 8601 format.
	RegisteredAt *time.Time `json:"registered_at,omitempty"`

	// Receipt's ID in online sales register. For successful receipt registration.
	FiscalProviderID string `json:"fiscal_provider_id,omitempty"`

	// List of products in the receipt (maximum 100 items).
	Items []ReceiptItem `json:"items,omitempty"`

	// List of completed payments.
	Settlements []yoocommon.Settlement `json:"settlements,omitempty"`

	// ID of the store on behalf of which you're sending the receipt.
	OnBehalfOf string `json:"on_behalf_of,omitempty"`

	// Store's taxation system.
	TaxSystemCode TaxSystemCode `json:"tax_system_code,omitempty"`

	// Industry attribute of the payment subject (tag 1260 in 54-FZ).
	ReceiptIndustryDetails *ReceiptIndustryDetails `json:"receipt_industry_details,omitempty"`

	// Transaction attribute of the receipt (tag 1270 in 54-FZ).
	ReceiptOperationalDetails *ReceiptOperationalDetails `json:"receipt_operational_details,omitempty"`
}
