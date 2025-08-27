package yoocommon

type Item struct {
	// parameter with the name of the product or service
	Description string `json:"description"`

	// parameter with the amount per unit of product
	Quantity string `json:"quantity"`

	// parameter specifying the quantity of goods (only integers, for example 1)
	Amount *Amount `json:"amount"`

	// parameter with the fixed value 1 (price without VAT)
	VatCode string `json:"vat_code"`

	// Payment method indicator (tag in Federal Law 54 — 1214) — reflects the type of payment and the fact of goods transfer.
	// Example: the customer fully pays for the goods and immediately receives them. In this case, the value full_payment (full settlement) must be passed.
	PaymentMode string `json:"payment_mode,omitempty"`

	// Item type indicator (tag in Federal Law 54 — 1212) — specifies what the payment is for, for example, a product or a service.
	PaymentSubject string `json:"payment_subject,omitempty"`
}
