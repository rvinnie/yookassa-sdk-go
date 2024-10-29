// Package yookassa implements all the necessary methods for working with YooMoney.
package yookassa

import (
	"encoding/json"
	"io"
	"net/http"

	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	yoosettings "github.com/rvinnie/yookassa-sdk-go/yookassa/settings"
)

const (
	MeEndpoint = "me"
)

// SettingsHandler works with client's account settings.
type SettingsHandler struct {
	client         *Client
	idempotencyKey string
}

func NewSettingsHandler(client *Client) *SettingsHandler {
	return &SettingsHandler{client: client}
}

func (r SettingsHandler) WithIdempotencyKey(idempotencyKey string) SettingsHandler {
	r.idempotencyKey = idempotencyKey

	return r
}

// GetAccountSettings gets the client account settings.
func (s *SettingsHandler) GetAccountSettings(OnBehalfOf *string) (*yoosettings.Settings, error) {
	var params map[string]interface{}
	if OnBehalfOf != nil {
		params = map[string]interface{}{"on_behalf_of": *OnBehalfOf}
	}

	resp, err := s.client.makeRequest(http.MethodGet, MeEndpoint, nil, params, s.idempotencyKey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	settingsResponse, err := s.parseSettingsResponse(resp)
	if err != nil {
		return nil, err
	}
	return settingsResponse, nil
}

func (s *SettingsHandler) parseSettingsResponse(
	resp *http.Response,
) (*yoosettings.Settings, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	settingsResponse := yoosettings.Settings{}
	err = json.Unmarshal(responseBytes, &settingsResponse)
	if err != nil {
		return nil, err
	}
	return &settingsResponse, nil
}
