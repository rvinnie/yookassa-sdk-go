package yoocommon

// Amount of Payment. Sometimes YooMoney's partners charge additional commission
// from the users that is not included in this amount.
type Amount struct {
	// Amount in the selected currency, in the form of a string with a dot separator,
	// for example, 10.00.
	Value string `json:"value,omitempty"`

	// Three-letter currency code in the ISO-4217 format. Example: RUB.
	Currency string `json:"currency,omitempty"`
}
