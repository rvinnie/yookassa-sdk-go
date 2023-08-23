package yooerrors

import (
	"encoding/json"
	"fmt"
	"io"
)

type YoomoneyError struct {
	Type        string `json:"type,omitempty"`
	Id          string `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Parameter   string `json:"parameter,omitempty"`
}

func (y *YoomoneyError) Error() string {
	return fmt.Sprintf("%s: %s", y.Type, y.Description)
}

func GetError(r io.Reader) (*YoomoneyError, error) {
	responseBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	errorBody := YoomoneyError{}
	err = json.Unmarshal(responseBytes, &errorBody)
	if err != nil {
		return &YoomoneyError{
			Type:        "error",
			Code:        "unexpected",
			Description: "Unexpected error occurs.",
		}, nil
	}

	return &errorBody, nil
}
