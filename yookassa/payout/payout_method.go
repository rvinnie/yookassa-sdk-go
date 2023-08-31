// Package yoopayout describes all the necessary entities for working with YooMoney Payouts.
package yoopayout

type PayoutMethodType string

const (
	// Payments to bank cards.
	PayoutTypeBankCard PayoutMethodType = "bank_card"
	// Payments to YooMoney wallets.
	PaymentTypeYooMoney PayoutMethodType = "yoo_money"
	// Payments via SBP.
	PayoutTypeSBP PayoutMethodType = "sbp"
)
