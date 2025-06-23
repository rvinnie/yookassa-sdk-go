package payout

import (
	"encoding/json"
	"reflect"
	"testing"

	yoopayout "github.com/imgrigorev/yookassa-sdk-go/yookassa/payout"
)

func TestPayoutBalance(t *testing.T) {
	payoutBalance := yoopayout.PayoutBalance{
		Value:    "320.00",
		Currency: "RUB",
	}

	jsonPayoutBalance, err := json.Marshal(payoutBalance)
	if err != nil {
		t.Error("Error in marshalling PayoutBalance object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonPayoutBalance, &got)
	if err != nil {
		t.Error("Error in unmarshalling PayoutBalance object")
	}

	expected := map[string]interface{}{
		"value":    "320.00",
		"currency": "RUB",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected PayoutBalance object are not equal")
	}
}
