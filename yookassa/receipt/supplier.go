// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Information about the supplier of product or service.
type Supplier struct {
	// Supplier name (tag 1225 in 54-FZ). The parameter is provided for by the format of fiscal documents (FFD)
	// and is considered mandatory for versions 1.1 and later.
	Name string `json:"name,omitempty"`

	// Supplier's phone number (tag 1171 in 54-FZ).
	// Specified in the ITU-T E.164 format, for example, 79000000000. The parameter is provided
	// for by the format of fiscal documents (FFD) and is considered mandatory for versions 1.1 and later.
	Phone string `json:"phone,omitempty"`

	// Provider's masked INN/TIN (tag 1226 in 54-FZ). Example: ***.
	// The parameter is provided for by the format of fiscal documents (FFD) and is considered mandatory for versions 1.05 and later.
	INN string `json:"inn,omitempty"`
}
