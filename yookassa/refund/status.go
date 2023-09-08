// Package yoorefund describes all the necessary entities for working with YooMoney Refunds.
package yoorefund

type RefundStatus string

const (
	// Pending - refund has been created, but it is still being processed.
	Pending RefundStatus = "pending"

	// Succeeded — refund successfully completed, the amount specified in the request
	// has been refunded to the user's means of payment (final and unchangeable status).
	Succeeded RefundStatus = "succeeded"

	// Canceled — refund canceled, the initiator and reasons are specified
	// in the cancellation_details object (final and unchangeable status).
	Canceled RefundStatus = "canceled"
)
