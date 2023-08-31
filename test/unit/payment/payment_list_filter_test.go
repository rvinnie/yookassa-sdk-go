package payment

import (
	"encoding/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"reflect"
	"testing"
)

func TestPaymentListFilter(t *testing.T) {
	paymentListFilter := yoopayment.PaymentListFilter{
		PaymentMethod: string(yoopayment.PaymentTypeBankCard),
		Status:        yoopayment.Succeeded,
		Cursor:        "37a5c87d-3984-51e8-a7f3-8de646d39ec15",
	}

	jsonPaymentListFilter, err := json.Marshal(paymentListFilter)
	if err != nil {
		t.Error("Error in marshalling PaymentListFilter object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonPaymentListFilter, &got)
	if err != nil {
		t.Error("Error in unmarshalling PaymentListFilter object")
	}

	expected := map[string]interface{}{
		"payment_method": "bank_card",
		"status":         "succeeded",
		"cursor":         "37a5c87d-3984-51e8-a7f3-8de646d39ec15",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected PaymentListFilter object are not equal")
	}
}
