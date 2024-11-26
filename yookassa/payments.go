// Package yookassa implements all the necessary methods for working with YooMoney.
package yookassa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

const (
	PaymentEndpoint = "payments"
	CaptureEndpoint = "capture"
	CancelEndpoint  = "cancel"
)

// PaymentHandler works with requests related to Payments.
type PaymentHandler struct {
	client         *Client
	idempotencyKey string
}

func NewPaymentHandler(client *Client) *PaymentHandler {
	return &PaymentHandler{client: client}
}

func (p PaymentHandler) WithIdempotencyKey(idempotencyKey string) *PaymentHandler {
	p.idempotencyKey = idempotencyKey

	return &p
}

// CapturePayment confirms payment, accepts and returns the Payment entity.
func (p *PaymentHandler) CapturePayment(payment *yoopayment.Payment) (*yoopayment.Payment, error) {
	paymentJson, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	captureRequest := fmt.Sprintf("%s/%s/%s", PaymentEndpoint, payment.ID, CaptureEndpoint)

	resp, err := p.client.makeRequest(
		http.MethodPost,
		captureRequest,
		paymentJson,
		nil,
		p.idempotencyKey,
	)
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

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}
	return paymentResponse, nil
}

// CancelPayment cancel payment by ID.
func (p *PaymentHandler) CancelPayment(paymentId string) (*yoopayment.Payment, error) {
	cancelRequest := fmt.Sprintf("%s/%s/%s", PaymentEndpoint, paymentId, CancelEndpoint)

	resp, err := p.client.makeRequest(http.MethodPost, cancelRequest, nil, nil, p.idempotencyKey)
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

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}
	return paymentResponse, nil
}

// CreatePayment creates a payment, accepts and returns the Payment entity.
func (p *PaymentHandler) CreatePayment(payment *yoopayment.Payment) (*yoopayment.Payment, error) {
	paymentJson, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(
		http.MethodPost,
		PaymentEndpoint,
		paymentJson,
		nil,
		p.idempotencyKey,
	)
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

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}

	if paymentResponse.Confirmation == nil {
		return nil, errors.New("empty confirmation url")
	}
	return paymentResponse, nil
}

// CreatePaymentLink creates a payment link, accepts Payment entity, returns the link.
func (p *PaymentHandler) CreatePaymentLink(payment *yoopayment.Payment) (string, error) {
	pay, err := p.CreatePayment(payment)
	if err != nil {
		return "", err
	}

	return p.ParsePaymentLink(pay)
}

// FindPayment find a payment by ID returns the Payment entity.
func (p *PaymentHandler) FindPayment(id string) (*yoopayment.Payment, error) {
	endpoint := fmt.Sprintf("%s/%s", PaymentEndpoint, id)

	resp, err := p.client.makeRequest(http.MethodGet, endpoint, nil, nil, p.idempotencyKey)
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

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}
	return paymentResponse, nil
}

// FindPayments find payments by filter and returns the list of payments.
func (p *PaymentHandler) FindPayments(
	filter *yoopayment.PaymentListFilter,
) (*yoopayment.PaymentList, error) {
	filterJson, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	var filterMap map[string]interface{}
	err = json.Unmarshal(filterJson, &filterMap)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(
		http.MethodGet,
		PaymentEndpoint,
		nil,
		filterMap,
		p.idempotencyKey,
	)
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

	var responseBytes []byte
	responseBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	paymentsResponse := yoopayment.PaymentList{}
	err = json.Unmarshal(responseBytes, &paymentsResponse)
	if err != nil {
		return nil, err
	}
	return &paymentsResponse, nil
}

// ParsePaymentLink retrieves a link to the Payment from the Payment entity.
func (p *PaymentHandler) ParsePaymentLink(payment *yoopayment.Payment) (string, error) {
	if payment == nil || payment.Confirmation == nil {
		return "", errors.New("empty confirmation url")
	}

	link, ok := payment.Confirmation.(map[string]interface{})["confirmation_url"].(string)
	if !ok {
		return "", errors.New("unable to get link")
	}
	return link, nil
}

func (p *PaymentHandler) parsePaymentResponse(resp *http.Response) (*yoopayment.Payment, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	paymentResponse := yoopayment.Payment{}
	err = json.Unmarshal(responseBytes, &paymentResponse)
	if err != nil {
		return nil, err
	}
	return &paymentResponse, nil
}
