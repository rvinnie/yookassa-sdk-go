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

func newClientWithPathAssertion(
	t *testing.T,
	expectedMethod string,
	expectedEscapedPath string,
	responseBody string,
) *Client {
	t.Helper()

	httpClient := http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			if req.Method != expectedMethod {
				t.Fatalf("unexpected method: %s", req.Method)
			}
			if req.URL.EscapedPath() != expectedEscapedPath {
				t.Fatalf("unexpected escaped path: %s", req.URL.EscapedPath())
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(responseBody)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	return NewClient("account_id", "secret_key", yooopts.WithHTTPClient(httpClient))
}

func TestCapturePaymentEscapesPaymentIDInPath(t *testing.T) {
	paymentID := "payment/id?test"
	client := newClientWithPathAssertion(
		t,
		http.MethodPost,
		"/v3/payments/payment%2Fid%3Ftest/capture",
		`{"id":"payment-id","status":"pending"}`,
	)
	handler := NewPaymentHandler(client)

	_, err := handler.CapturePayment(context.Background(), &yoopayment.Payment{ID: paymentID})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestCancelPaymentEscapesPaymentIDInPath(t *testing.T) {
	paymentID := "payment/id?test"
	client := newClientWithPathAssertion(
		t,
		http.MethodPost,
		"/v3/payments/payment%2Fid%3Ftest/cancel",
		`{"id":"payment-id","status":"pending"}`,
	)
	handler := NewPaymentHandler(client)

	_, err := handler.CancelPayment(context.Background(), paymentID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestFindPaymentEscapesPaymentIDInPath(t *testing.T) {
	paymentID := "payment/id?test"
	client := newClientWithPathAssertion(
		t,
		http.MethodGet,
		"/v3/payments/payment%2Fid%3Ftest",
		`{"id":"payment-id","status":"pending"}`,
	)
	handler := NewPaymentHandler(client)

	_, err := handler.FindPayment(context.Background(), paymentID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestFindRefundEscapesRefundIDInPath(t *testing.T) {
	refundID := "refund/id?test"
	client := newClientWithPathAssertion(
		t,
		http.MethodGet,
		"/v3/refunds/refund%2Fid%3Ftest",
		`{"id":"refund-id","status":"pending"}`,
	)
	handler := NewRefundHandler(client)

	_, err := handler.FindRefund(context.Background(), refundID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestGetPayoutEscapesPayoutIDInPath(t *testing.T) {
	payoutID := "payout/id?test"
	client := newClientWithPathAssertion(
		t,
		http.MethodGet,
		"/v3/payouts/payout%2Fid%3Ftest",
		`{"id":"payout-id","status":"pending"}`,
	)
	handler := NewPayoutHandler(client)

	_, err := handler.GetPayout(context.Background(), payoutID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}
