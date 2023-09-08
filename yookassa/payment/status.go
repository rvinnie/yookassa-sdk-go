// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

type PaymentStatus string

const (
	// Pending - data is being processed.
	Pending PaymentStatus = "pending"

	// Waiting for capture.
	WaitingForCapture PaymentStatus = "waiting_for_capture"

	// Succeeded — receipt successfully registered.
	Succeeded PaymentStatus = "succeeded"

	// Canceled — receipt was not registered, you need to create it independently.
	Canceled PaymentStatus = "canceled"
)
