// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// The payment subject attribute is specified in the payment_subject parameter
// https://yookassa.ru/developers/payment-acceptance/receipts/54fz/parameters-values#payment-subject
type PaymentSubject string

const (
	// Available starting from FFD 1.05

	Commodity            PaymentSubject = "commodity"
	Excise               PaymentSubject = "excise"
	Job                  PaymentSubject = "job"
	Service              PaymentSubject = "service"
	GamblingBet          PaymentSubject = "gambling_bet"
	GamblingPrize        PaymentSubject = "gambling_prize"
	Lottery              PaymentSubject = "lottery"
	LotteryPrize         PaymentSubject = "lottery_prize"
	IntellectualActivity PaymentSubject = "intellectual_activity"
	Payment              PaymentSubject = "payment"
	AgentCommission      PaymentSubject = "agent_commission"
	PropertyRight        PaymentSubject = "property_right"
	NonOperatingGain     PaymentSubject = "non_operating_gain"
	InsurancePremium     PaymentSubject = "insurance_premium"
	SalesTax             PaymentSubject = "sales_tax"
	ResortFee            PaymentSubject = "resort_fee"
	Composite            PaymentSubject = "composite"
	Another              PaymentSubject = "another"

	// Available starting from FFD 1.2

	Marked                         PaymentSubject = "marked"
	NonMarked                      PaymentSubject = "non_marked"
	MarkedExcise                   PaymentSubject = "marked_excise"
	NonMarkedExcise                PaymentSubject = "non_marked_excise"
	Fine                           PaymentSubject = "fine"
	Tax                            PaymentSubject = "tax"
	Lien                           PaymentSubject = "lien"
	Cost                           PaymentSubject = "cost"
	AgentWithdrawals               PaymentSubject = "agent_withdrawals"
	PensionInsuranceWithoutPayouts PaymentSubject = "pension_insurance_without_payouts"
	PensionInsuranceWithPayouts    PaymentSubject = "pension_insurance_with_payouts"
	HealthInsuranceWithoutPayouts  PaymentSubject = "health_insurance_without_payouts"
	HealthInsuranceWithPayouts     PaymentSubject = "health_insurance_with_payouts"
	HealthInsurance                PaymentSubject = "health_insurance"
)
