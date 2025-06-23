package settings

import (
	"encoding/json"
	"reflect"
	"testing"

	yoosettings "github.com/imgrigorev/yookassa-sdk-go/yookassa/settings"
)

func TestPayoutBalance(t *testing.T) {
	status := yoosettings.Enabled

	settings := yoosettings.Settings{
		AccountId:            "123",
		Status:               &status,
		FiscalizationEnabled: true,
		ITN:                  "1111111111",
		Name:                 "hello",
	}

	jsonSettings, err := json.Marshal(settings)
	if err != nil {
		t.Error("Error in marshalling Settings object")
	}

	var got map[string]interface{}
	err = json.Unmarshal(jsonSettings, &got)
	if err != nil {
		t.Error("Error in unmarshalling Settings object")
	}

	expected := map[string]interface{}{
		"account_id":            "123",
		"fiscalization_enabled": true,
		"itn":                   "1111111111",
		"name":                  "hello",
		"status":                "enabled",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got and expected Settings object are not equal")
	}
}
