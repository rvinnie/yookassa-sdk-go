// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

// Banking details of the payer (legal entity or sole proprietor).
type PayerBankDetails struct {
	// Full name of the organization.
	FullName string `json:"full_name,omitempty" binding:"max=800"`

	// Abbreviated name of the organization.
	ShortName string `json:"short_name,omitempty" binding:"max=160"`

	// Address of the organization.
	Address string `json:"address,omitempty" binding:"max=500"`

	// Taxpayer Identification Number (INN) of the organization.
	INN string `json:"inn,omitempty"`

	// Name of the organization's bank.
	BankName string `json:"bank_name,omitempty" binding:"min=1,max=350"`

	// Branch of the organization's bank.
	BankBranch string `json:"bank_branch,omitempty" binding:"min=1,max=140"`

	// Bank Identification Code (BIC) of the organization's bank.
	BankBIK string `json:"bank_bik,omitempty"`

	// Account number of the organization.
	Account string `json:"account,omitempty"`

	// Tax Registration Reason Code (KPP) of the organization.
	KPP string `json:"kpp,omitempty"`
}
