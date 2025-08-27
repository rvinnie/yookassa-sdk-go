package yoocommon

type Customer struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`

	// For a legal entity — the name of the organization; for an individual entrepreneur or a private individual — full name. If an individual does not have a TIN, passport details should be provided in this parameter instead. No more than 256 characters.
	// This can be provided if you use receipts from YooKassa or online cash registers Orange Data, Atol Online.
	FullName string `json:"full_name,omitempty"`

	// User’s TIN (10 or 12 digits). If an individual does not have a TIN, passport details must be passed in the full_name parameter.
	// This can be provided if you use receipts from YooKassa or online cash registers Orange Data, Atol Online.
	INN string `json:"inn,omitempty"`
}
