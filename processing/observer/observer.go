package observer

import (
	"context"
	"github.com/HucciK/crypto-processing/core"
	"github.com/HucciK/crypto-processing/errors"
	"github.com/HucciK/crypto-processing/processing/config"
	"time"
)

type Client interface {
	Call(params core.CallParams) (*core.CallResult, error)
}

type ObserveFunc func(ctx context.Context, w core.Wallet, amount uint64) (chan struct{}, error)

type ObserverService struct {
	config    config.ObserverConfig
	client    Client
	observers []ObserveFunc
}

func NewObserverService(config config.ObserverConfig, client Client) *ObserverService {
	o := &ObserverService{
		config: config,
		client: client,
	}
	o.observers = []ObserveFunc{
		o.tronObserver,
	}

	return o
}

func (o *ObserverService) Observe(ctx context.Context, wallet core.Wallet, amount uint64) (chan struct{}, error) {
	for _, obs := range o.observers {
		ch, err := obs(ctx, wallet, amount)
		if err != nil {
			return nil, err
		}

		if ch != nil {
			return ch, nil
		}
	}

	return nil, errors.ErrCantCreateObserver
}

func (o *ObserverService) tronObserver(ctx context.Context, wallet core.Wallet, amount uint64) (chan struct{}, error) {
	if wallet.GetChain() != core.ChainTron {
		return nil, nil
	}
	observer := make(chan struct{})

	params := core.CallParams{
		Chain:        wallet.GetChain(),
		Method:       core.CallMethodTransactionList,
		Wallet:       wallet,
		ContractAddr: "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	}

	go func() {
		defer close(observer)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.NewTimer(time.Duration(o.config.ObserveInterval) * time.Millisecond).C:
				res, err := o.client.Call(params)
				if err != nil {
					continue
				}

				success, err := o.checkForSuccess(res.TransactionList, amount, params.ContractAddr)
				if err != nil {
					continue
				}

				if !success {
					continue
				}

				observer <- struct{}{}
				return
			}
		}
	}()

	return observer, nil
}

func (o *ObserverService) checkForSuccess(list []core.Transaction, amount uint64, contract string) (bool, error) {
	for _, trx := range list {
		if trx.GetTimestamp() > time.Now().Add(-15*time.Minute).Unix() && trx.IsSuccessful() && trx.GetContractAddress() == contract && trx.GetAmount() >= amount {
			return true, nil
		}
	}
	return false, nil
}
