package yoowebhook

import yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"

type WebhookType string

const (
	WebhookTypeNotification WebhookType = "notification"
)

type WebhookEventType string

const (
	EventPaymentSucceeded         WebhookEventType = "payment.succeeded"
	EventPaymentWaitingForCapture WebhookEventType = "payment.waiting_for_capture"
	EventPaymentCanceled          WebhookEventType = "payment.canceled"
	EventRefundSucceeded          WebhookEventType = "refund.succeeded"
)

type WebhookEvent struct {
	Type   WebhookType        `json:"type"`
	Event  WebhookEventType   `json:"event"`
	Object yoopayment.Payment `json:"object"`
}
