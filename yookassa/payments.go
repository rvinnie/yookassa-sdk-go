package yookassa

import (
	"encoding/json"
	"errors"
	"fmt"
	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"io"
	"net/http"
)

const (
	PaymentEndpoint = "payments"
	CaptureEndpoint = "capture"
	CancelEndpoint  = "cancel"
)

type PaymentHandler struct {
	client *Client
}

func NewPaymentHandler(client *Client) *PaymentHandler {
	return &PaymentHandler{client: client}
}

func (p *PaymentHandler) CapturePayment(payment *yoopayment.Payment) (*yoopayment.Payment, error) {
	paymentJson, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	captureRequest := fmt.Sprintf("%s/%s/%s", PaymentEndpoint, payment.ID, CaptureEndpoint)
	resp, err := p.client.makeRequest(http.MethodPost, captureRequest, paymentJson, nil)
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

func (p *PaymentHandler) CancelPayment(paymentId string) (*yoopayment.Payment, error) {
	cancelRequest := fmt.Sprintf("%s/%s/%s", PaymentEndpoint, paymentId, CancelEndpoint)
	fmt.Println(cancelRequest)
	resp, err := p.client.makeRequest(http.MethodPost, cancelRequest, nil, nil)
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

func (p *PaymentHandler) CreatePayment(payment *yoopayment.Payment) (*yoopayment.Payment, error) {
	paymentJson, err := json.MarshalIndent(payment, "", "\t")
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(http.MethodPost, PaymentEndpoint, paymentJson, nil)
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

func (p *PaymentHandler) CreatePaymentLink(payment *yoopayment.Payment) (string, error) {
	pay, err := p.CreatePayment(payment)
	if err != nil {
		return "", err
	}

	return p.ParsePaymentLink(pay)
}

func (p *PaymentHandler) FindPayment(id string) (*yoopayment.Payment, error) {
	endpoint := fmt.Sprintf("%s/%s", PaymentEndpoint, id)
	resp, err := p.client.makeRequest(http.MethodGet, endpoint, nil, nil)
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

func (p *PaymentHandler) FindPayments(filter *yoopayment.PaymentListFilter) (*yoopayment.PaymentList, error) {
	filterJson, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	var filterMap map[string]interface{}
	err = json.Unmarshal(filterJson, &filterMap)
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest(http.MethodGet, PaymentEndpoint, nil, filterMap)
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
