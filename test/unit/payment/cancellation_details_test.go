package payment

import (
	"encoding/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	"reflect"
	"testing"
)

func TestCancellationDetails(t *testing.T) {
	cancellationDetails := yoopayment.CancellationDetails{
		Party:  "yoo_money",
		Reason: "deal_expired",
	}

	jsonCancellationDetails, err := json.Marshal(cancellationDetails)
	if err != nil {
		t.Error("Error in marshalling CancellationDetails object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonCancellationDetails, &got)
	if err != nil {
		t.Error("Error in unmarshalling CancellationDetails object")
	}

	expected := map[string]interface{}{
		"party":  "yoo_money",
		"reason": "deal_expired",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected CancellationDetails object are not equal")
	}
}
