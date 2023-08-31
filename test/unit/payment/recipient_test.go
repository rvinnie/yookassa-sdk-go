package payment

import (
	"encoding/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"reflect"
	"testing"
)

func TestRecipient(t *testing.T) {
	recipient := yoopayment.Recipient{
		AccountId: "123",
		GatewayId: "456",
	}

	jsonRecipient, err := json.Marshal(recipient)
	if err != nil {
		t.Error("Error in marshalling Recipient object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonRecipient, &got)
	if err != nil {
		t.Error("Error in unmarshalling Recipient object")
	}

	expected := map[string]interface{}{
		"account_id": "123",
		"gateway_id": "456",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected Recipient object are not equal")
	}
}
