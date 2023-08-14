package payment

import "time"

// The Payment object contains all currently relevant information
// about the payment. The object is generated during creation of a payment,
// and sent in response to any payment-related requests.
type Payment struct {
	// Payment ID in YooMoney.
	ID string `json:"id" binding:"required"`

	// Payment Status. Possible values: pending, waiting_for_capture, succeeded, and canceled.
	Status Status `json:"status" binding:"required"`

	// Payment Amount. Sometimes YooMoney's partners charge additional
	// commission from the users that is not included in this amount.
	Amount Amount `json:"amount" binding:"required"`

	// Amount of payment to be received by the store: the amount value minus the YooMoney commission.
	IncomeAmount Amount `json:"income_amount"`

	// Description of the transaction (maximum 128 characters) displayed in your YooMoney
	// Merchant Profile, and shown to the user during checkout. For example,
	// "Payment for order No. 72 for user@yoomoney.ru".
	Description string `json:"description" binding:"max=128"`

	// Payment Recipient.
	Recipient Recipient `json:"recipient" binding:"required"`

	// Payment method used for this payment.
	PaymentMethod PaymentMethoder `json:"payment_method"`

	// Time of order creation, based on UTC and specified in the ISO 8601 format.
	// Example: 2017-11-03T11:52:31.827Z
	CapturedAt time.Time `json:"captured_at"`

	// Time of order creation, based on UTC and specified in the ISO 8601 format.
	// Example: 2017-11-03T11:52:31.827Z
	CreatedAt time.Time `json:"created_at" binding:"required"`

	// The period during which you can cancel or capture a payment for free.
	// The payment with the waiting_for_capture status will be automatically
	// canceled at the specified time. Based on UTC and specified in the ISO 8601 format.
	// Example: 2017-11-03T11:52:31.827Z
	ExpiresAt time.Time `json:"expires_at"`

	// Selected payment confirmation scenario.
	// For payments requiring confirmation from the user.
	Confirmation Confirmer `json:"confirmation"`

	// The attribute of a test transaction.
	Test bool `json:"test" binding:"required"`

	// The amount refunded to the user. Specified if the payment has successful refunds.
	RefundedAmount Amount `json:"refunded_amount"`

	// The attribute of a paid order.
	Paid bool `json:"paid" binding:"required"`

	// Availability of the option to make a refund via API.
	Refundable bool `json:"refundable" binding:"required"`

	// Status of receipt delivery.
	ReceiptRegistration Status `json:"receipt_registration"`

	// Any additional data you might require for processing payments
	// (for example, your internal order ID), specified as a “key-value” pair and
	// returned in response from YooMoney. Limitations: no more than 16 keys,
	// no more than 32 characters in the key’s title, no more than 512 characters
	// in the key’s value, data type is a string in the UTF-8 format.
	Metadata interface{} `json:"metadata"`

	// Commentary to the canceled status: who and why canceled the payment.
	CancellationDetails CancellationDetails `json:"cancellation_details"`

	// Payment authorization details.
	AuthorizationDetails AuthorizationDetails `json:"authorization_details"`

	// Information about money distribution: the amounts of transfers and
	// the stores to be transferred to.
	Transfers []Transfer `json:"transfers"`

	// The deal within which the payment is being carried out.
	Deal Deal `json:"deal"`

	// The identifier of the customer in your system, such as email address or phone number.
	// No more than 200 characters.
	MerchantCustomerID string `json:"merchant_customer_id" binding:"max=200"`
}
