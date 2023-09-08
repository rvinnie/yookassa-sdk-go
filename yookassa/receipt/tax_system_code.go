// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Store's taxation system.
type TaxSystemCode int

const (
	// General tax system.
	GeneralTaxSystem TaxSystemCode = 1

	// Simplified (STS, income).
	SimplifiedIncome TaxSystemCode = 2

	// Simplified (STS, income with costs deducted)
	SimplifiedIncomeWithCosts TaxSystemCode = 3

	// Unified tax on imputed income (ENVD)
	ENVD TaxSystemCode = 4

	// Unified agricultural tax (ESN)
	ESN TaxSystemCode = 5

	// Patent Based Tax System
	PatentBasedTaxSystem TaxSystemCode = 6
)
