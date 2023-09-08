// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Fraction of a marked product (tag 1291 in 54-FZ).
type MarkQuantity struct {
	// The number of products sold from one customer package (tag 1293 in 54-FZ).
	// Cannot exceed the denominator.
	Numerator int64 `json:"numerator,omitempty"`

	// The total number of products in the customer package (tag 1294 in 54-FZ).
	Denominator int64 `json:"denominator,omitempty"`
}
