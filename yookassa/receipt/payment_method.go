// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// The payment method attribute is specified in the payment_mode parameter.
type PaymentMethod string

const (
	FullPrepayment    PaymentMethod = "full_prepayment"
	PartialPrepayment PaymentMethod = "partial_prepayment"
	Advance           PaymentMethod = "advance"
	FullPayment       PaymentMethod = "full_payment"
	PartialPayment    PaymentMethod = "partial_payment"
	Credit            PaymentMethod = "credit"
	CreditPayment     PaymentMethod = "credit_payment"
)
