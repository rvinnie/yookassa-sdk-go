package payment

// The Deal within which the payment is being carried out.
type Deal struct {
	// Deal ID.
	ID string `json:"id" binding:"required,min=36,max=50"`

	// Information about money distribution.
	Settlements []Settlement `json:"settlements" binding:"required"`
}
