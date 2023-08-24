package yookassa

import (
	"encoding/json"
	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/settings"
	"io"
	"net/http"
)

const (
	MeEndpoint = "me"
)

type SettingsHandler struct {
	client *Client
}

func NewSettingsHandler(client *Client) *SettingsHandler {
	return &SettingsHandler{client: client}
}

func (s *SettingsHandler) GetAccountSettings(OnBehalfOf *string) (*yoosettings.Settings, error) {
	var params map[string]interface{}
	if OnBehalfOf != nil {
		params = map[string]interface{}{"on_behalf_of": *OnBehalfOf}
	}

	resp, err := s.client.makeRequest(http.MethodGet, MeEndpoint, nil, params)
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

func (s *SettingsHandler) parseSettingsResponse(resp *http.Response) (*yoosettings.Settings, error) {
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
