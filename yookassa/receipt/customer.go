// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// User details.
type Customer struct {
	// Name of the organization for companies, full name for sole proprietors and individuals.
	// If the individual doesn't have a Tax Identification Number (INN),
	// specify their passport information in this parameter. Maximum 256 characters.
	// Online sales register that support this parameter: Orange Data, ATOL Online.
	FullName string `json:"full_name,omitempty"`

	// User's Tax Identification Number (INN) (10 or 12 digits).
	// If the individual doesn't have an INN, specify their passport information in the full_name parameter.
	// Online sales register that support this parameter: Orange Data, ATOL Online.
	INN string `json:"inn,omitempty"`

	// User's email address for sending the receipt. Required parameter if phone isn't specified.
	Email string `json:"email,omitempty"`

	// User's phone number for sending the receipt.
	// Specified in the ITU-T E.164 format, for example, 79000000000. Required parameter if email isn't specified.
	Phone string `json:"phone,omitempty"`
}
