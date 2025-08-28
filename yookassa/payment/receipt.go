package yoopayment

import yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"

type Receipt struct {
	// Customer info
	Customer *yoocommon.Customer `json:"customer"`

	// Array with the data regarding products or services (maximum 6 positions)
	Items []*yoocommon.Item `json:"items"`

	// For third-party online cash registers: a required parameter if you use an Atol Online cash register updated to FFD 1.2, or if you have multiple taxation systems; in all other cases, it is not passed.
	TaxSystemCode int16 `json:"tax_system_code,omitempty"`
}
