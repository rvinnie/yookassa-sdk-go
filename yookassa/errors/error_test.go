package yooerrors

import (
	"strings"
	"testing"
)

func TestGetErrorLimitsResponseBodySize(t *testing.T) {
	payload := strings.Repeat(" ", maxErrorBodyBytes+1) +
		`{"type":"error","id":"error-id","code":"invalid_request","description":"bad request"}`

	result, err := GetError(strings.NewReader(payload))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Code != "unexpected" {
		t.Fatalf("expected fallback code for oversized body, got: %s", result.Code)
	}
}
