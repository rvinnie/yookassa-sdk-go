// Package yookassa implements all the necessary methods for working with YooMoney.
package yookassa

import (
	"encoding/json"
	"fmt"
	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	yooreceipt "github.com/rvinnie/yookassa-sdk-go/yookassa/receipt"
	"io"
	"net/http"
)

const (
	ReceiptEndpoint = "receipts"
)

// ReceiptHandler works with requests related to Receipts.
type ReceiptHandler struct {
	client *Client
}

func NewReceiptHandler(client *Client) *ReceiptHandler {
	return &ReceiptHandler{client: client}
}

// CreateReceipt creates a receipt, accepts and returns the Receipt entity.
func (r *ReceiptHandler) CreateReceipt(receipt *yooreceipt.ReceiptInput) (*yooreceipt.Receipt, error) {
	receiptJson, err := json.Marshal(receipt)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.makeRequest(http.MethodPost, ReceiptEndpoint, receiptJson, nil)
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

	receiptResponse, err := r.parseReceiptResponse(resp)
	if err != nil {
		return nil, err
	}

	return receiptResponse, nil
}

// FindReceipt find a receipt by ID returns the Receipt entity.
func (r *ReceiptHandler) FindReceipt(id string) (*yooreceipt.Receipt, error) {
	endpoint := fmt.Sprintf("%s/%s", ReceiptEndpoint, id)
	resp, err := r.client.makeRequest(http.MethodGet, endpoint, nil, nil)
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

	receiptResponse, err := r.parseReceiptResponse(resp)
	if err != nil {
		return nil, err
	}
	return receiptResponse, nil
}

// FindReceipts find receipts by filter and returns the list of receipts.
func (r *ReceiptHandler) FindReceipts(filter *yooreceipt.ReceiptListFilter) (*yooreceipt.ReceiptList, error) {
	filterJson, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	var filterMap map[string]interface{}
	err = json.Unmarshal(filterJson, &filterMap)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.makeRequest(http.MethodGet, ReceiptEndpoint, nil, filterMap)
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

	receiptsResponse := yooreceipt.ReceiptList{}
	err = json.Unmarshal(responseBytes, &receiptsResponse)
	if err != nil {
		return nil, err
	}
	return &receiptsResponse, nil
}

func (r *ReceiptHandler) parseReceiptResponse(resp *http.Response) (*yooreceipt.Receipt, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	receiptResponse := yooreceipt.Receipt{}
	err = json.Unmarshal(responseBytes, &receiptResponse)
	if err != nil {
		return nil, err
	}
	return &receiptResponse, nil
}
