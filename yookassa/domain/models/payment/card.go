package payment

// Card details.
type Card struct {
	// First 6 digits of the card’s number (BIN).
	// For payments with bank cards saved in YooMoney and other services,
	// the specified BIN might not correspond with the last4, expiry_year, expiry_month values.
	First6 string `json:"first6"`

	// Last 4 digits of the card's number.
	Last4 string `json:"last4" binding:"required"`

	// Expiration date, year, YYYY.
	ExpiryYear string `json:"expiry_year" binding:"required"`

	// Expiration date, month, MM.
	ExpiryMonth string `json:"expiry_month" binding:"required"`

	// Type of bank card. Possible values: MasterCard (for Mastercard and Maestro cards),
	// Visa (for Visa and Visa Electron cards), Mir, UnionPay, JCB, AmericanExpress,
	// DinersClub, DiscoverCard, InstaPayment, InstaPaymentTM, Laser, Dankort,
	// Solo, Switch, and Unknown.
	CardType string `json:"card_type" binding:"required"`

	// Code of the country where the bank card was issued according to ISO-3166 alpha-2.
	// Example: RU.
	IssuerCountry string `json:"issuer_country"`

	// Name of the issuing bank.
	IssuerName string `json:"issuer_name"`

	// Source of bank card details. Possible values: mir_pay, apple_pay, google_pay.
	// For payments where the user selects a card saved in Mir Pay, Apple Pay or Google Pay.
	Source string `json:"source"`
}
