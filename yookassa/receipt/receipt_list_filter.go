// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import "time"

// The request allows you to receive the list of receipts filtered by specified criteria.
type ReceiptListFilter struct {
	// Filter by creation date: time must be greater than the specified
	// value or equal ("from a certain moment inclusive").
	// Specified in the ISO 8601 format.
	CreatedAtGTE *time.Time `json:"created_at.gte,omitempty"`

	// Filter by creation date: time must be greater than the specified
	// value ("from a certain moment exclusive").
	// Specified in the ISO 8601 format.
	CreatedAtGT *time.Time `json:"created_at.gt,omitempty"`

	// Filter by creation date: time must be less than the specified
	// value or equal ("until a certain moment inclusive").
	// Specified in the ISO 8601 format
	CreatedAtLTE *time.Time `json:"created_at.lte,omitempty"`

	// Filter by creation date: time must be less than the specified
	// value ("until a certain moment exclusive").
	// Specified in the ISO 8601 format.
	CreatedAtLT *time.Time `json:"created_at.lt,omitempty"`

	// Filter by receipt status. Example: succeeded
	Status ReceiptStatus `json:"status,omitempty"`

	// Filter by payment ID.
	PaymentID string `json:"payment_id,omitempty"`

	// Filter by refund ID.
	RefundID string `json:"refund_id,omitempty"`

	// Size of the output of request results: number of objects sent in response.
	// Possible values: 1 to 100.
	// Default value: 10
	Limit int `json:"limit,omitempty"`

	// Cursor to the next fragment in the list. Example: cursor=37a5c87d-3984-51e8-a7f3-8de646d39ec15
	// Use the value of the next_cursor parameter received in response to the previous request as the cursor.
	// Used if the size of the list is greater than the output size (limit) and the output end hasn't been reached.
	Cursor string `json:"cursor,omitempty"`
}
