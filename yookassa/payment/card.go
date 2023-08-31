// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

// Card details.
type Card struct {
	// First 6 digits of the card’s number (BIN).
	// For payments with bank cards saved in YooMoney and other services,
	// the specified BIN might not correspond with the last4, expiry_year, expiry_month values.
	First6 string `json:"first6,omitempty"`

	// Last 4 digits of the card's number.
	Last4 string `json:"last4,omitempty"`

	// Expiration date, year, YYYY.
	ExpiryYear string `json:"expiry_year,omitempty"`

	// Expiration date, month, MM.
	ExpiryMonth string `json:"expiry_month,omitempty"`

	// Type of bank card. Possible values: MasterCard (for Mastercard and Maestro cards),
	// Visa (for Visa and Visa Electron cards), Mir, UnionPay, JCB, AmericanExpress,
	// DinersClub, DiscoverCard, InstaPayment, InstaPaymentTM, Laser, Dankort,
	// Solo, Switch, and Unknown.
	CardType string `json:"card_type,omitempty"`

	// Code of the country where the bank card was issued according to ISO-3166 alpha-2.
	// Example: RU.
	IssuerCountry string `json:"issuer_country,omitempty"`

	// Name of the issuing bank.
	IssuerName string `json:"issuer_name,omitempty"`

	// Source of bank card details. Possible values: mir_pay, apple_pay, google_pay.
	// For payments where the user selects a card saved in Mir Pay, Apple Pay or Google Pay.
	Source string `json:"source,omitempty"`
}
