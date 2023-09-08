// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

import yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"

// Product in the receipt.
type ReceiptItem struct {
	// Product name (maximum 128 characters). Tag 1030 in 54-FZ.
	Description string `json:"description,omitempty"`

	// Product quantity. Tag 1023 in 54-FZ.
	// Maximum possible value depends on the model of your online sales register.
	// Format: decimal number, with the fractional part of three or more characters
	// (number depends on quantity in the request). Fractional part separated by the decimal point.
	Quantity float64 `json:"quantity,omitempty"`

	// Product price.
	Amount *yoocommon.Amount `json:"amount,omitempty"`

	// VAT rate. Possible value is a number from 1 to 6.
	VATCode VATCode `json:"vat_code,omitempty"`

	// Payment subject attribute (tag 1212 in 54-FZ): what the payment is made for, for example, a product or service.
	PaymentSubject PaymentSubject `json:"payment_subject,omitempty"`

	// Payment method attribute (tag 1214 in 54-FZ): contains information about the payment method
	// and shows whether the product has been handed over to the customer.
	PaymentMode PaymentMethod `json:"payment_mode,omitempty"`

	// Country of origin code according to the Russian classifier of world countries (OK (MK (ISO 3166) 004-97) 025-2001).
	// Example: RU.
	// Online sales register that support this parameter: Orange Data, Kit Invest.
	CountryOfOriginCode string `json:"country_of_origin_code,omitempty"`

	// Customs declaration number (1 to 32 characters).
	// Online sales register that support this parameter: Orange Data, Kit Invest.
	CustomsDeclarationNumber string `json:"customs_declaration_number,omitempty"`

	// Amount of excise tax on products including kopeks. Decimal number with 2 digits after the period.
	// Online sales register that support this parameter: Orange Data, Kit Invest.
	Excise string `json:"excise,omitempty"`

	// Information about the supplier of product or service (tag 1224 in 54-FZ).
	// You can specify this parameter if you send the data for creating the receipt using the receipt after payment scenario.
	Supplier *Supplier `json:"supplier,omitempty"`

	// Type of agent selling goods or services.
	// The parameter is provided for by the format of fiscal documents (FFD) and is considered mandatory for versions 1.1 and later.
	AgentType AgentType `json:"agent_type,omitempty"`

	// Product code (tag 1163 in 54-FZ).
	// Must be specified if the FFD 1.2 protocol is used and if the product must be marked.
	// At least one of the fields must be filled in.
	MarkCodeInfo *MarkCodeInfo `json:"mark_code_info,omitempty"`

	// Unit of measurement of product quantity: for example, items or grams. Tag 2108 in 54-FZ.
	// This parameter must be specified starting from FFD 1.2.
	Measure Measure `json:"measure,omitempty"`

	// Industry attribute of the payment subject (tag 1260 in 54-FZ). Must be specified if FFD 1.2 is used.
	PaymentSubjectIndustryDetails *ReceiptOperationalDetails `json:"payment_subject_industry_details,omitempty"`

	// Product code is a unique number assigned to a unit of product during marking process (tag 1162 in 54-FZ).
	// Format: hexadecimal number with spaces. Maximum length is 32 bytes.
	ProductCode string `json:"product_code,omitempty"`

	//Method of processing the marking code (tag 2102 in 54-FZ). Must be specified if all of the following applies:
	//
	//- FFD version 1.2 is used;
	//
	//- payment is made for a marked product;
	//
	//- an online sales register from ATOL Online or BusinessRu is used.
	//
	//Must get a value equal to "0".
	MarkMode string `json:"mark_mode,omitempty"`

	// Fraction of a marked product (tag 1291 in 54-FZ). Must be specified if all of the following applies:
	//
	//- FFD version 1.2 is used;
	//
	//- payment is made for a marked product;
	//
	//- the measure field has the piece value.
	//
	// Example: you're selling pencils by the piece.
	// They're supplied in packages 100 pencils each with one marking code.
	// To sell one pencil, enter 1 in numerator and 100 in denominator.
	MarkQuantity *MarkQuantity `json:"mark_quantity,omitempty"`
}
