// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// VAT rate code is specified in the vat_code parameter. Possible value is a number from 1 to 6.
type VATCode int

const (
	// VAT not included.
	NotIncluded VATCode = 1

	// 0% VAT rate.
	VAT0 VATCode = 2

	// 10% VAT rate.
	VAT10 VATCode = 3

	// 20% receipt’s VAT rate.
	VAT20 VATCode = 4

	// 10/110 receipt’s estimate VAT rate.
	VAT10110 VATCode = 5

	// 20/120 receipt’s estimate VAT rate.
	VAT20120 VATCode = 6
)
