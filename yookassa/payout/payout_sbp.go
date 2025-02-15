package yoopayout

type SbpBank struct {
	BankId string `json:"bank_id"`
	Name   string `json:"name"`
	Bic    string `json:"bic"`
}

type SbpBankList struct {
	Type  string    `json:"type"` // possible value: list
	Items []SbpBank `json:"items"`
}
