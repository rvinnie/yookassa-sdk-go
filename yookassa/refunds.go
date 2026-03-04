// Package yookassa implements all the necessary methods for working with YooMoney.
package yookassa

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	yoorefund "github.com/rvinnie/yookassa-sdk-go/yookassa/refund"
)

const (
	RefundEndpoint = "refunds"
)

// RefundHandler works with requests related to Refunds.
type RefundHandler struct {
	client         Requester
	idempotencyKey string
}

func NewRefundHandler(client Requester) *RefundHandler {
	return &RefundHandler{client: client}
}

func (r RefundHandler) WithIdempotencyKey(idempotencyKey string) RefundHandler {
	r.idempotencyKey = idempotencyKey

	return r
}

// CreateRefund creates a refund, accepts and returns the Refund entity.
func (r *RefundHandler) CreateRefund(ctx context.Context, refund *yoorefund.Refund) (*yoorefund.Refund, error) {
	refundJson, err := json.Marshal(refund)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.MakeRequest(
		ctx,
		http.MethodPost,
		RefundEndpoint,
		refundJson,
		nil,
		r.idempotencyKey,
	)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer func() { _ = resp.Body.Close() }()
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	refundResponse, err := r.parseRefundResponse(resp)
	if err != nil {
		return nil, err
	}

	return refundResponse, nil
}

// FindRefund find a refund by ID returns the Refund entity.
func (r *RefundHandler) FindRefund(ctx context.Context, id string) (*yoorefund.Refund, error) {
	endpoint := fmt.Sprintf("%s/%s", RefundEndpoint, url.PathEscape(id))

	resp, err := r.client.MakeRequest(ctx, http.MethodGet, endpoint, nil, nil, r.idempotencyKey)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer func() { _ = resp.Body.Close() }()
	}

	if resp.StatusCode != http.StatusOK {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	refundResponse, err := r.parseRefundResponse(resp)
	if err != nil {
		return nil, err
	}
	return refundResponse, nil
}

// FindRefunds find refunds by filter and returns the list of refunds.
func (r *RefundHandler) FindRefunds(
	ctx context.Context,
	filter *yoorefund.RefundListFilter,
) (*yoorefund.RefundList, error) {
	filterJson, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	var filterMap map[string]interface{}
	err = json.Unmarshal(filterJson, &filterMap)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.MakeRequest(
		ctx,
		http.MethodGet,
		RefundEndpoint,
		nil,
		filterMap,
		r.idempotencyKey,
	)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer func() { _ = resp.Body.Close() }()
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
	responseBytes, err = io.ReadAll(io.LimitReader(resp.Body, maxResponseBodyBytes))
	if err != nil {
		return nil, err
	}

	refundsResponse := yoorefund.RefundList{}
	err = json.Unmarshal(responseBytes, &refundsResponse)
	if err != nil {
		return nil, err
	}
	return &refundsResponse, nil
}

func (r *RefundHandler) parseRefundResponse(resp *http.Response) (*yoorefund.Refund, error) {
	var responseBytes []byte
	responseBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseBodyBytes))
	if err != nil {
		return nil, err
	}

	refundResponse := yoorefund.Refund{}
	err = json.Unmarshal(responseBytes, &refundResponse)
	if err != nil {
		return nil, err
	}
	return &refundResponse, nil
}
