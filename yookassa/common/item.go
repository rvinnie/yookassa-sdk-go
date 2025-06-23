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

	// Признак способа расчета full_prepayment || full_payment
	PaymentMode string `json:"payment_mode"`

	// Признак предмета расчета
	PaymentSubject string `json:"payment_subject"`
}
