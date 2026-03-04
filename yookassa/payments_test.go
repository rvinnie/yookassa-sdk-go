package yookassa

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	yooopts "github.com/rvinnie/yookassa-sdk-go/yookassa/opts"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type trackingReadCloser struct {
	io.Reader
	closed bool
}

func (t *trackingReadCloser) Close() error {
	t.closed = true

	return nil
}

func newPaymentHandlerWithRoundTrip(rt roundTripFunc) *PaymentHandler {
	httpClient := http.Client{Transport: rt}
	client := NewClient("account_id", "secret_key", yooopts.WithHTTPClient(httpClient))

	return NewPaymentHandler(client)
}

func newPaymentHandlerWithResponse(statusCode int, body string) *PaymentHandler {
	return newPaymentHandlerWithRoundTrip(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	}))
}

func TestCreatePaymentClosesResponseBodyOnSuccess(t *testing.T) {
	var responseBody *trackingReadCloser
	handler := newPaymentHandlerWithRoundTrip(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		responseBody = &trackingReadCloser{
			Reader: strings.NewReader(
				`{"id":"payment-id","status":"pending","confirmation":{"type":"redirect","confirmation_url":"https://example.com"}}`,
			),
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       responseBody,
			Header:     make(http.Header),
		}, nil
	}))

	_, err := handler.CreatePayment(context.Background(), &yoopayment.Payment{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if responseBody == nil || !responseBody.closed {
		t.Fatal("expected response body to be closed")
	}
}

func TestCreatePaymentClosesResponseBodyOnError(t *testing.T) {
	var responseBody *trackingReadCloser
	handler := newPaymentHandlerWithRoundTrip(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		responseBody = &trackingReadCloser{
			Reader: strings.NewReader(
				`{"type":"error","id":"error-id","code":"invalid_request","description":"bad request"}`,
			),
		}

		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       responseBody,
			Header:     make(http.Header),
		}, nil
	}))

	_, err := handler.CreatePayment(context.Background(), &yoopayment.Payment{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if responseBody == nil || !responseBody.closed {
		t.Fatal("expected response body to be closed")
	}
}

func TestCreatePaymentReturnsErrorWhenBodyExceedsLimit(t *testing.T) {
	oversizedResponse := strings.Repeat(" ", maxResponseBodyBytes+1) +
		`{"id":"payment-id","status":"pending","confirmation":{"type":"redirect","confirmation_url":"https://example.com"}}`
	handler := newPaymentHandlerWithResponse(http.StatusOK, oversizedResponse)

	_, err := handler.CreatePayment(context.Background(), &yoopayment.Payment{})
	if err == nil {
		t.Fatal("expected error for oversized response body, got nil")
	}
}

func TestParsePaymentLinkReturnsErrorForUnexpectedConfirmationType(t *testing.T) {
	handler := NewPaymentHandler(nil)

	_, err := handler.ParsePaymentLink(&yoopayment.Payment{
		Confirmation: yoopayment.Redirect{
			Type:            yoopayment.TypeRedirect,
			ConfirmationURL: "https://example.com",
		},
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "unable to get link" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreatePaymentReturnsErrorWhenConfirmationIsEmptyAndPaymentMethodIDIsEmpty(t *testing.T) {
	handler := newPaymentHandlerWithResponse(http.StatusOK, `{"id":"payment-id","status":"pending"}`)

	_, err := handler.CreatePayment(context.Background(), &yoopayment.Payment{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "empty confirmation url" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreatePaymentDoesNotReturnErrorWhenConfirmationIsEmptyAndPaymentMethodIDIsSet(t *testing.T) {
	handler := newPaymentHandlerWithResponse(http.StatusOK, `{"id":"payment-id","status":"pending"}`)

	result, err := handler.CreatePayment(context.Background(), &yoopayment.Payment{PaymentMethodID: "pm-123"})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result == nil {
		t.Fatal("expected payment response, got nil")
	}

	if result.ID != "payment-id" {
		t.Fatalf("unexpected payment id: %s", result.ID)
	}
}
