package payment

import (
	"encoding/json"
	"reflect"
	"testing"

	yoopayment "github.com/imgrigorev/yookassa-sdk-go/yookassa/payment"
)

func TestCard(t *testing.T) {
	creditCard := yoopayment.Card{
		First6:        "555555",
		Last4:         "4444",
		ExpiryYear:    "2022",
		ExpiryMonth:   "07",
		CardType:      "MasterCard",
		IssuerCountry: "RU",
		IssuerName:    "Sberbank",
	}

	jsonCard, err := json.Marshal(creditCard)
	if err != nil {
		t.Error("Error in marshalling Card object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonCard, &got)
	if err != nil {
		t.Error("Error in unmarshalling Card object")
	}

	expected := map[string]interface{}{
		"first6":         "555555",
		"last4":          "4444",
		"expiry_month":   "07",
		"expiry_year":    "2022",
		"card_type":      "MasterCard",
		"issuer_country": "RU",
		"issuer_name":    "Sberbank",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected credit Card object are not equal")
	}
}
