package payment

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
)

func TestIdempotencyKey(t *testing.T) {
	idempotencyKey := uuid.NewString()

	paymentHandler := yookassa.NewPaymentHandler(nil)

	paymentHandler.WithIdempotencyKey(idempotencyKey)

	if paymentHandler.IdempotencyKey() == idempotencyKey {
		t.Errorf("Wrong behaviour of idempotency key: %s", idempotencyKey)
	}

	if paymentHandler.IdempotencyKey() != "" {
		t.Errorf("Idempotency key must be set only for one request")
	}
}
