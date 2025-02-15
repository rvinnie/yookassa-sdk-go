package yoopayout

import yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"

// See https://yookassa.ru/developers/payouts/making-payouts/sbp?lang=en

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type PayoutDestinationData struct {
	Type   PayoutMethodType `json:"type"`
	Phone  string           `json:"phone"`
	BankId string           `json:"bank_id"`
}

type Metadata struct {
	OrderId string `json:"order_id"`
}

type Payout struct {
	Id                    string                `json:"id"`
	Amount                *yoocommon.Amount     `json:"amount"`
	PayoutDestinationData PayoutDestinationData `json:"payout_destination_data"`
	Description           string                `json:"description"`
	Metadata              Metadata              `json:"metadata"`
	CreatedAt             string                `json:"created_at"`
	Status                Status                `json:"status"`
	Test                  bool                  `json:"test"`
}
