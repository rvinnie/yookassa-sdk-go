// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Product code (tag 1163 in 54-FZ).
// Must be specified if the FFD 1.2 protocol is used and if the product must be marked
type MarkCodeInfo struct {
	// Product code as it was read by the scanner (tag 2000 in 54-FZ).
	// Must be specified if an online sales register from Orange Data is used.
	MarkCodeRaw string `json:"mark_code_raw,omitempty"`

	// Unknown product code (tag 1300 in 54-FZ).
	Unknown string `json:"unknown,omitempty"`

	// Product code in the EAN-8 format (tag 1301 in 54-FZ).
	EAN8 string `json:"ean_8,omitempty"`

	// Product code in the EAN-13 format (tag 1302 in 54-FZ).
	EAN13 string `json:"ean_13,omitempty"`

	// Product code in the ITF-14 format (tag 1303 in 54-FZ).
	ITF14 string `json:"itf_14,omitempty"`

	// Product code in the GS1.0 format (tag 1304 in 54-FZ).
	// Online sales register that support this parameter: Orange Data, aQsi, Kit Invest.
	GS10 string `json:"gs_10,omitempty"`

	// Product code in the GS1.M format (tag 1305 in 54-FZ).
	GS1M string `json:"gs_1m,omitempty"`

	// Product code in the short marking code format (tag 1306 in 54-FZ).
	Short string `json:"short,omitempty"`

	// Control ID of a fur product (tag 1307 in 54-FZ).
	Fur string `json:"fur"`

	// Product code in the EGAIS-2.0 (unified state automated information system) format. Tag 1308 in 54-FZ.
	EGAIS20 string `json:"egais_20,omitempty"`

	// Product code in the EGAIS-3.0 (unified state automated information system) format. Tag 1309 in 54-FZ.
	EGAIS30 string `json:"egais_30,omitempty"`
}
