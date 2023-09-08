// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Status of receipt delivery.
type ReceiptStatus string

const (
	// Pending - data is being processed.
	Pending ReceiptStatus = "pending"

	// Succeeded — receipt successfully registered.
	Succeeded ReceiptStatus = "succeeded"

	// Canceled — receipt was not registered.
	Canceled ReceiptStatus = "canceled"
)
