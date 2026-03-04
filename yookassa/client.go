// Package yookassa implements all the necessary methods for working with YooMoney.
package yookassa

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	yooopts "github.com/rvinnie/yookassa-sdk-go/yookassa/opts"
)

const (
	BaseURL              = "https://api.yookassa.ru/v3/"
	defaultHTTPTimeout   = 30 * time.Second
	maxResponseBodyBytes = 10 << 20 // 10 MiB
)

// HTTPDoer is an abstraction over http.Client used by SDK Client.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Requester is an abstraction used by handlers to perform API requests.
type Requester interface {
	MakeRequest(
		ctx context.Context,
		method string,
		endpoint string,
		body []byte,
		params map[string]interface{},
		idempotencyKey string,
	) (*http.Response, error)
}

// Client works with YooMoney API.
type Client struct {
	client    HTTPDoer
	accountId string
	secretKey string
}

func NewClient(accountId string, secretKey string, options ...yooopts.Option) *Client {
	config := yooopts.Config{}
	for _, option := range options {
		if option != nil {
			option.Apply(&config)
		}
	}

	httpClient := config.Client
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	return &Client{
		client:    httpClient,
		accountId: accountId,
		secretKey: secretKey,
	}
}

func (c *Client) MakeRequest(
	ctx context.Context,
	method string,
	endpoint string,
	body []byte,
	params map[string]interface{},
	idempotencyKey string,
) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	uri := fmt.Sprintf("%s%s", BaseURL, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if idempotencyKey == "" {
		idempotencyKey = uuid.NewString()
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotence-Key", idempotencyKey)
	}

	req.SetBasicAuth(c.accountId, c.secretKey)

	if params != nil {
		q := req.URL.Query()
		for paramName, paramVal := range params {
			q.Add(paramName, fmt.Sprintf("%v", paramVal))
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
