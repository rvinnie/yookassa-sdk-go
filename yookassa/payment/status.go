// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

type Status string

const (
	// Pending - data is being processed.
	Pending Status = "pending"

	// Waiting for capture.
	WaitingForCapture Status = "waiting_for_capture"

	// Succeeded — receipt successfully registered.
	Succeeded Status = "succeeded"

	// Canceled — receipt was not registered, you need to create it independently.
	Canceled Status = "canceled"
)
