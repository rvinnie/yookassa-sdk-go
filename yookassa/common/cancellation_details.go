package yoocommon

// Commentary to the canceled status: who and why canceled the payment.
type CancellationDetails struct {
	// The participant of the payment process that
	// made the decision to cancel the payment. Possible values are yoo_money,
	// payment_network, and merchant.
	Party string `json:"party,omitempty"`

	// Reason behind the cancelation.
	Reason string `json:"reason,omitempty"`
}
