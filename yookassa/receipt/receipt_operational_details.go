// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import "time"

// Transaction attribute of the receipt (tag 1270 in 54-FZ).
type ReceiptOperationalDetails struct {
	// Transaction ID (tag 1271 in 54-FZ). From 0 to 255 characters.
	OperationID string `json:"operation_id,omitempty"`

	// Transaction details (tag 1272 in 54-FZ).
	Value string `json:"value,omitempty"`

	// Time when the transaction was initiated (tag 1273 in 54-FZ).
	// Formatted in accordance with UTC standart and specified in the ISO 8601.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
