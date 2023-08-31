package payment

import (
	"encoding/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"reflect"
	"testing"
)

func TestPayerBankDetails(t *testing.T) {
	payerBankDetails := yoopayment.PayerBankDetails{
		FullName:   "Limited liability company 'Organization'",
		ShortName:  "LLC 'Organization'",
		Address:    "197111, Russian Federation, St. Petersburg, 3rd Severovokzalny st., 17, building/building 2, apt. 16",
		INN:        "7728662610",
		BankName:   "SBERBANK",
		BankBranch: "NORTH-WESTERN BANK OF SBERBANK RF",
		BankBIK:    "9449899",
		Account:    "40702810355002140000",
		KPP:        "783501610",
	}

	jsonPayerBankDetails, err := json.Marshal(payerBankDetails)
	if err != nil {
		t.Error("Error in marshalling PayerBankDetails object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonPayerBankDetails, &got)
	if err != nil {
		t.Error("Error in unmarshalling PayerBankDetails object")
	}

	expected := map[string]interface{}{
		"full_name":   "Limited liability company 'Organization'",
		"short_name":  "LLC 'Organization'",
		"address":     "197111, Russian Federation, St. Petersburg, 3rd Severovokzalny st., 17, building/building 2, apt. 16",
		"inn":         "7728662610",
		"bank_name":   "SBERBANK",
		"bank_branch": "NORTH-WESTERN BANK OF SBERBANK RF",
		"bank_bik":    "9449899",
		"account":     "40702810355002140000",
		"kpp":         "783501610",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected PayerBankDetails object are not equal")
	}
}
