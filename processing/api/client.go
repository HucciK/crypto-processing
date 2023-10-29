package api

import (
	"encoding/json"
	"fmt"
	"github.com/HucciK/crypto-processing/core"
	"github.com/HucciK/crypto-processing/errors"
	"github.com/HucciK/crypto-processing/processing/crypto/tron"
	"io"
	"net/http"
)

type CallFunc func(params core.CallParams) (*core.CallResult, error)

type Client struct {
	client  http.Client
	apis    []CallFunc
	methods map[core.Chain][]CallFunc
}

func NewClient() *Client {
	c := &Client{
		client: http.Client{},
	}

	c.apis = []CallFunc{
		c.tronAPICall,
	}

	c.methods = map[core.Chain][]CallFunc{
		core.ChainTron: {c.tronTransactionList},
	}

	return c
}

func (c *Client) Call(params core.CallParams) (*core.CallResult, error) {
	for _, call := range c.apis {
		res, err := call(params)
		if err != nil {
			return nil, err
		}

		if res != nil {
			return res, nil
		}
	}

	return nil, errors.ErrEmptyCallResult
}

func (c *Client) tronAPICall(params core.CallParams) (*core.CallResult, error) {
	if params.Chain != core.ChainTron {
		return nil, nil
	}

	for _, m := range c.methods[params.Chain] {
		res, err := m(params)
		if err != nil {
			return nil, err
		}

		if res != nil {
			return res, nil
		}
	}

	return nil, errors.ErrEmptyCallResult
}

func (c *Client) tronTransactionList(params core.CallParams) (*core.CallResult, error) {
	url := fmt.Sprintf("https://api.shasta.trongrid.io/v1/accounts/%s/transactions", params.Wallet.GetAddress())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tronTransactions tron.TronTrasactions
	if err := json.Unmarshal(data, &tronTransactions); err != nil {
		return nil, err
	}

	if tronTransactions.Transactions == nil {
		return nil, errors.ErrNoTransactionFound
	}

	return &core.CallResult{
		TransactionList: tronTransactions.ToList(),
	}, nil
}
