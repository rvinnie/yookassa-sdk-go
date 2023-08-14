package payment

// Banking details of the payer (legal entity or sole proprietor).
type PayerBankDetails struct {
	// Full name of the organization.
	FullName string `json:"full_name" binding:"required,max=800"`

	// Abbreviated name of the organization.
	ShortName string `json:"short_name" binding:"required,max=160"`

	// Address of the organization.
	Address string `json:"address" binding:"required,max=500"`

	// Taxpayer Identification Number (INN) of the organization.
	INN string `json:"inn" binding:"required"`

	// Name of the organization's bank.
	BankName string `json:"bank_name" binding:"required,min=1,max=350"`

	// Branch of the organization's bank.
	BankBranch string `json:"bank_branch" binding:"required,min=1,max=140"`

	// Bank Identification Code (BIC) of the organization's bank.
	BankBIK string `json:"bank_bik" binding:"required"`

	// Account number of the organization.
	Account string `json:"account" binding:"required"`

	// Tax Registration Reason Code (KPP) of the organization.
	KPP string `json:"kpp"`
}
