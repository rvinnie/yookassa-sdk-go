package payment

// Payment authorization details.
type AuthorizationDetails struct {
	// Retrieval Reference Number is a unique identifier of a transaction
	// in the issuer's system. Used for payments via bank card.
	RRN string `json:"rrn"`

	// Bank card's authorization code.
	// Provided by the issuer to confirm authorization.
	AuthCode string `json:"auth_code"`

	// Information about user’s 3‑D Secure authentication for confirming the payment.
	ThreeDSecure struct {
		// Information on whether the 3-D Secure authentication form
		// is displayed to the user for confirming the payment or not
		Applied bool `json:"applied"`
	} `json:"three_d_secure"`
}
