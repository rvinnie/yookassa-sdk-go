// Package yoorefund describes all the necessary entities for working with YooMoney Refunds.
package yoorefund

type Status string

const (
	// Pending - refund has been created, but it is still being processed.
	Pending Status = "pending"

	// Succeeded — refund successfully completed, the amount specified in the request
	// has been refunded to the user's means of payment (final and unchangeable status).
	Succeeded Status = "succeeded"

	// Canceled — refund canceled, the initiator and reasons are specified
	// in the cancellation_details object (final and unchangeable status).
	Canceled Status = "canceled"
)
