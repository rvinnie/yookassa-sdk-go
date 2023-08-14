package payment

// Information about money distribution:
// the amounts of transfers and the stores to be transferred to.
type Transfer struct {
	// ID of the store in favor of which you're accepting the receipt.
	// Provided by YooMoney, displayed in the Sellers section of your Merchant Profile (shopId column).
	AccountID string `json:"account_id" binding:"required"`

	// Amount to be transferred to the store.
	Amount Amount `json:"amount" binding:"required"`

	// Status of the money distribution between stores.
	Status Status `json:"status" binding:"required"`

	// Commission for sold products or services charged in your favor.
	PlatformFeeAmount Amount `json:"platform_fee_amount"`

	// Transaction description, which the seller will see in the YooMoney Merchant Profile.
	// Example: "Marketplace order No. 72".
	Description string `json:"status" binding:"max=128"`

	// Any additional data you might require for processing payments
	// (for example, your internal order ID), specified as a “key-value” pair and returned in response from YooMoney.
	// Limitations: no more than 16 keys, no more than 32 characters in the key’s title,
	// no more than 512 characters in the key’s value, data type is a string in the UTF-8 format.
	Metadata interface{} `json:"metadata"`
}
