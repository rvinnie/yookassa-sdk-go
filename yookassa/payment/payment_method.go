// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

type PaymentMethodType string

const (
	PaymentTypeAlfabank      PaymentMethodType = "alfabank"
	PaymentTypeApplePay      PaymentMethodType = "apple_pay"
	PaymentTypeB2BSberbank   PaymentMethodType = "b2b_sberbank"
	PaymentTypeBankCard      PaymentMethodType = "bank_card"
	PaymentTypeCash          PaymentMethodType = "cash"
	PaymentTypeGooglePay     PaymentMethodType = "google_pay"
	PaymentTypeInstallments  PaymentMethodType = "installments"
	PaymentTypeMobileBalance PaymentMethodType = "mobile_balance"
	PaymentTypeQiwi          PaymentMethodType = "qiwi"
	PaymentTypeSberbank      PaymentMethodType = "sberbank"
	PaymentTypeSBP           PaymentMethodType = "sbp"
	PaymentTypeTinkoffBank   PaymentMethodType = "tinkoff_bank"
	PaymentTypeWebmoney      PaymentMethodType = "webmoney"
	PaymentTypeWeChat        PaymentMethodType = "wechat"
	PaymentTypeYooMoney      PaymentMethodType = "yoo_money"
)

type PaymentMethoder interface {
}

type paymentMethod struct {
	// Payment method code.
	Type PaymentMethodType `json:"type,omitempty"`

	// Payment method ID.
	ID string `json:"id,omitempty"`

	// Saving payment methods allows conducting automatic recurring payments.
	Saved bool `json:"saved,omitempty"`

	// Payment method name.
	Title string `json:"title,omitempty"`
}

type Alfabank struct {
	paymentMethod

	// User's login in Alfa-Click (linked phone number or the additional login).
	Login string `login:"login,omitempty"`
}

type ApplePay struct {
	paymentMethod
}

type B2BSberbank struct {
	paymentMethod

	// Banking details of the payer (legal entity or sole proprietor).
	PayerBankDetails PayerBankDetails `json:"payer_bank_details,omitempty"`

	// Purpose of payment (no more than 210 characters).
	PaymentPurpose string `json:"payment_purpose,omitempty" binding:"max=210"`

	// Information about the value-added tax (VAT).
	// A payment might or might not be subject to VAT.
	// Products may be taxed at the same VAT rate, or at different rates.
	VATData string `json:"vat_data,omitempty"`
}

type BankCard struct {
	paymentMethod

	// Bank card details.
	Card Card `json:"card,omitempty"`
}

type Cash struct {
	paymentMethod
}

type GooglePay struct {
	paymentMethod
}

type Installments struct {
	paymentMethod
}

type MobileBalance struct {
	paymentMethod
}

type Qiwi struct {
	paymentMethod
}

type Sberbank struct {
	paymentMethod

	// Bank card details.
	Card Card `json:"card,omitempty"`

	// The phone number specified during the registration process
	// of the SberPay account, specified in the ITU-T E.164 format,
	// for example, 79000000000.
	Phone string `json:"phone,omitempty"`
}

type SBP struct {
	paymentMethod
}

type TinkoffBank struct {
	paymentMethod
}

type WebMoney struct {
	paymentMethod
}

type WeChat struct {
	paymentMethod
}

type YooMoney struct {
	paymentMethod

	// The number of the YooMoney wallet used for making the payment.
	AccountNumber string `json:"account_number,omitempty"`
}
