// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Additional user parameter
type AdditionalUserProps struct {
	// Name of the additional user parameter (tag 1085 in 54-FZ). No more than 64 characters.
	Name string `json:"name,omitempty"`

	//Value of the additional user parameter (tag 1086 in 54-FZ). No more than 234 characters.
	Value string `json:"value,omitempty"`
}
