package payment

type Recipient struct {
	// Store's ID in YooMoney.
	AccountId string `json:"account_id,omitempty"`

	// Subaccount's ID. Used for separating payment flows within one account.
	GatewayId string `json:"gateway_id,omitempty"`
}
