// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

type ConfirmationType string

const (
	TypeEmbedded          ConfirmationType = "embedded"
	TypeExternal          ConfirmationType = "external"
	TypeMobileApplication ConfirmationType = "mobile_application"
	TypeQR                ConfirmationType = "qr"
	TypeRedirect          ConfirmationType = "redirect"
)

type Confirmer interface {
}

type Embedded struct {
	// Confirmation scenario code.
	Type ConfirmationType `json:"type,omitempty"`

	// Token for the YooMoney Checkout Widget initialization.
	ConfirmationToken string `json:"confirmation_token,omitempty"`
}

type External struct {
	// Confirmation scenario code.
	Type ConfirmationType `json:"type,omitempty"`
}

type MobileApplication struct {
	// Confirmation scenario code.
	Type ConfirmationType `json:"type,omitempty"`

	// Deep link to the mobile app where the user confirms the payment.
	ConfirmationURL string `json:"confirmation_url,omitempty"`
}

type QR struct {
	// Confirmation scenario code.
	Type ConfirmationType `json:"type,omitempty"`

	// Data for generating the QR code.
	ConfirmationData string `json:"confirmation_data,omitempty"`
}

type Redirect struct {
	// Confirmation scenario code.
	Type ConfirmationType `json:"type,omitempty"`

	// The URL that the user will be redirected to for payment confirmation.
	ConfirmationURL string `json:"confirmation_url,omitempty"`

	// A request for making a payment with authentication by 3-D Secure.
	// It works if you accept bank card payments without user confirmation by default.
	// In other cases, the 3-D Secure authentication will be handled by YooMoney.
	// If you would like to accept payments without additional confirmation by the user,
	// contact your YooMoney manager.
	Enforce bool `json:"enforce,omitempty"`

	// The URL that the user will return to after confirming or
	// canceling the payment on the webpage.
	ReturnURL string `json:"return_url,omitempty" binding:"max=2048"`
}
