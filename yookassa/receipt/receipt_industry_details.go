// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import "time"

// Industry attribute of the payment subject (tag 1260 in 54-FZ).
type ReceiptIndustryDetails struct {
	// ID of the federal executive authority (tag 1262 in 54-FZ).
	FederalID string `json:"federal_id,omitempty"`

	// Date of the incorporation document. Tag 1263 in 54-FZ. Specified in the ISO 8601 format.
	DocumentDate *time.Time `json:"document_date,omitempty"`

	// Number of the regulation issued by the federal executive authority prescribing
	// how the "Industry attribute value" attribute must be filled in. Tag 1264 in 54-FZ.
	DocumentNumber string `json:"document_number,omitempty"`

	// Industry attribute value (tag 1265 in 54-FZ).
	Value string `json:"value,omitempty"`
}
