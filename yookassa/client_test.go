package yookassa

import (
	"net/http"
	"testing"
	"time"

	yooopts "github.com/rvinnie/yookassa-sdk-go/yookassa/opts"
)

func TestNewClientSetsDefaultTimeout(t *testing.T) {
	client := NewClient("account_id", "secret_key")
	httpClient, ok := client.client.(*http.Client)
	if !ok {
		t.Fatalf("unexpected client type: %T", client.client)
	}

	if httpClient.Timeout != defaultHTTPTimeout {
		t.Fatalf("unexpected default timeout: %s", httpClient.Timeout)
	}

	if httpClient.Timeout <= 0 {
		t.Fatal("expected positive default timeout")
	}
}

func TestNewClientUsesProvidedHTTPClient(t *testing.T) {
	custom := http.Client{Timeout: 5 * time.Second}
	client := NewClient("account_id", "secret_key", yooopts.WithHTTPClient(custom))
	httpClient, ok := client.client.(*http.Client)
	if !ok {
		t.Fatalf("unexpected client type: %T", client.client)
	}

	if httpClient.Timeout != custom.Timeout {
		t.Fatalf("unexpected custom timeout: %s", httpClient.Timeout)
	}
}
