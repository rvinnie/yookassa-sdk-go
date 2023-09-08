// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Type of receipt in the online sales register.
type ReceiptType string

const (
	PaymentType ReceiptType = "payment"
	RefundType  ReceiptType = "refund"
)
