package yookassa

import (
	"encoding/json"
	"errors"

	yooerror "github.com/rvinnie/yookassa-sdk-go/yookassa/errors"
	yoopayout "github.com/rvinnie/yookassa-sdk-go/yookassa/payout"
)

const (
	SbpBanksEndpoint = "sbp_banks"
	PayoutsEndpoint  = "payouts"
)

// PayoutHandler works with requests related to Payouts.
type PayoutHandler struct {
	client         *Client
	idempotencyKey string
}

func NewPayoutHandler(client *Client) *PayoutHandler {
	return &PayoutHandler{client: client}
}

func (p PayoutHandler) WithIdempotencyKey(idempotencyKey string) *PayoutHandler {
	p.idempotencyKey = idempotencyKey

	return &p
}

func (p *PayoutHandler) GetSbpBanks() ([]yoopayout.SbpBank, error) {
	resp, err := p.client.makeRequest("GET", SbpBanksEndpoint, nil, nil, p.idempotencyKey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	var sbpBanks yoopayout.SbpBankList
	err = json.NewDecoder(resp.Body).Decode(&sbpBanks)
	if err != nil {
		return nil, err
	}

	return sbpBanks.Items, nil
}

// TODO: support other payout types
func (p *PayoutHandler) CreatePayout(payout *yoopayout.Payout) (*yoopayout.Payout, error) {
	if payout.PayoutDestinationData.Type != yoopayout.PayoutTypeSBP {
		return nil, errors.New("unsupported payout type")
	}

	payoutJson, err := json.MarshalIndent(payout, "", "\t")
	if err != nil {
		return nil, err
	}

	resp, err := p.client.makeRequest("POST", PayoutsEndpoint, payoutJson, nil, p.idempotencyKey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	var createdPayout yoopayout.Payout
	err = json.NewDecoder(resp.Body).Decode(&createdPayout)
	if err != nil {
		return nil, err
	}

	return &createdPayout, nil
}

func (p *PayoutHandler) GetPayout(payoutId string) (*yoopayout.Payout, error) {
	endpoint := PayoutsEndpoint + "/" + payoutId
	resp, err := p.client.makeRequest("GET", endpoint, nil, nil, p.idempotencyKey)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var respError error
		respError, err = yooerror.GetError(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, respError
	}

	var payout yoopayout.Payout
	err = json.NewDecoder(resp.Body).Decode(&payout)
	if err != nil {
		return nil, err
	}

	return &payout, nil
}
