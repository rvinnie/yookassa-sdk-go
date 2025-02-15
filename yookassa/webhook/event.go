package yoowebhook

import (
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	yoopayout "github.com/rvinnie/yookassa-sdk-go/yookassa/payout"
)

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
	EventPayoutSucceeded          WebhookEventType = "payout.succeeded"
	EventPayoutCanceled           WebhookEventType = "payout.canceled"
)

type WebhookEvent[T yoopayment.Payment | yoopayout.Payout] struct {
	Type   WebhookType      `json:"type"`
	Event  WebhookEventType `json:"event"`
	Object T                `json:"object"`
}
