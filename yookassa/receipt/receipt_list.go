// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// ReceiptList contains list of receipts. Data can be filtered.
type ReceiptList struct {
	// Format of request results output. Possible value: list.
	Type string `json:"type"`

	// Array of objects matching the request parameters.
	Items []Receipt `json:"items"`

	// Cursor to the next fragment of the list.
	// Example: cursor=37a5c87d-3984-51e8-a7f3-8de646d39ec15
	NextCursor string `json:"next_cursor"`
}
