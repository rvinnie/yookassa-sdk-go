package yookassa

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/domain/models/payment"
	"io"
	"net/http"
)

type PaymentHandler struct {
	client *Client
}

func NewPaymentHandler(client *Client) *PaymentHandler {
	return &PaymentHandler{client: client}
}

func (p *PaymentHandler) CreatePayment(payment payment.Payment) (*payment.Payment, error) {
	paymentJson, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(http.MethodPost, "payments", paymentJson)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var responseBytes []byte
		responseBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(responseBytes))
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

func (p *PaymentHandler) CreatePaymentLink(payment payment.Payment) (string, error) {
	pay, err := p.CreatePayment(payment)
	if err != nil {
		return "", err
	}

	return p.ParsePaymentLink(pay)
}

func (p *PaymentHandler) FindPayment(id string) (*payment.Payment, error) {
	resp, err := p.client.makeRequest(http.MethodGet, "payments/"+id, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		var responseBytes []byte
		responseBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(responseBytes))
	}

	paymentResponse, err := p.parsePaymentResponse(resp)
	if err != nil {
		return nil, err
	}
	return paymentResponse, nil
}

func (p *PaymentHandler) ParsePaymentLink(payment *payment.Payment) (string, error) {
	if payment == nil || payment.Confirmation == nil {
		return "", errors.New("empty confirmation url")
	}

	link, ok := payment.Confirmation.(map[string]interface{})["confirmation_url"].(string)
	if !ok {
		return "", errors.New("unable to get link")
	}
	return link, nil
}

func (p *PaymentHandler) parsePaymentResponse(resp *http.Response) (*payment.Payment, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	paymentResponse := payment.Payment{}
	err = json.Unmarshal(responseBytes, &paymentResponse)
	if err != nil {
		return nil, err
	}
	return &paymentResponse, nil
}
