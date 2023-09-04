package payment

import (
	"encoding/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	"reflect"
	"testing"
)

func TestAmount(t *testing.T) {
	amount := yoopayment.Amount{
		Value:    "1.97",
		Currency: "RUB",
	}

	jsonAmount, err := json.Marshal(amount)
	if err != nil {
		t.Error("Error in marshalling Amount object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonAmount, &got)
	if err != nil {
		t.Error("Error in unmarshalling Amount object")
	}

	expected := map[string]interface{}{
		"value":    "1.97",
		"currency": "RUB",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected Amount object are not equal")
	}
}
