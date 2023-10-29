package observer

import (
	"context"
	"github.com/HucciK/crypto-processing/internal/core"
	"github.com/HucciK/crypto-processing/internal/processing/config"
	"github.com/HucciK/crypto-processing/internal/processing/crypto/tron"
	mock_observer "github.com/HucciK/crypto-processing/internal/processing/observer/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestObserverService_Observe_EmptyTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	observeInterval := 10
	ctxTimeout := 20
	waitInterval := 30

	cli := mock_observer.NewMockClient(ctrl)
	obs := NewObserverService(config.ObserverConfig{
		ObserveInterval: observeInterval,
	}, cli)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Millisecond)
	w := tron.TronWallet{
		Chain:      core.ChainTron,
		PrivateKey: "",
		Address:    "TEd6j8KG8qdZW1uESMBv233AzZXkAEPB1P",
		Balance:    0,
	}

	params := core.CallParams{
		Chain:        w.GetChain(),
		Method:       core.CallMethodTransactionList,
		Wallet:       &w,
		ContractAddr: "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	}

	expResp := &core.CallResult{
		TransactionList: []core.Transaction{
			&tron.TronTransacion{},
			&tron.TronTransacion{},
			&tron.TronTransacion{},
		},
	}

	cli.EXPECT().Call(params).Return(expResp, nil).AnyTimes()
	ch, err := obs.Observe(ctx, &w, 10)
	require.Nil(t, err)
	require.NotNil(t, ch)

	time.Sleep(time.Duration(waitInterval) * time.Millisecond)
	v, ok := <-ch
	require.Empty(t, v)
	require.Equal(t, false, ok)
}

func TestObserverService_Observe_ValidTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	observeInterval := 10
	ctxTimeout := 20
	waitInterval := 30

	cli := mock_observer.NewMockClient(ctrl)
	obs := NewObserverService(config.ObserverConfig{
		ObserveInterval: observeInterval,
	}, cli)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Millisecond)
	w := tron.TronWallet{
		Chain:      core.ChainTron,
		PrivateKey: "",
		Address:    "TEd6j8KG8qdZW1uESMBv233AzZXkAEPB1P",
		Balance:    0,
	}

	params := core.CallParams{
		Chain:        w.GetChain(),
		Method:       core.CallMethodTransactionList,
		Wallet:       &w,
		ContractAddr: "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	}

	var list tron.TronTrasactions
	list.Transactions = []tron.TronTransacion{
		{
			Ret:  []tron.Ret{{ContractRet: tron.TronTransactionStatusSuccess}},
			TxID: "",
			RawData: tron.RawData{
				Contract: []tron.Contract{{
					Parameter: tron.ContractParameter{Value: tron.ParameterValue{
						Data:            "a9059cbb000000000000000000000000076abaa36a89610946e77273a704e2231c3897200000000000000000000000000000000000000000000000000000000000a7d8c0",
						Amount:          0,
						OwnerAddress:    "",
						ToAddress:       "",
						ContractAddress: "TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
					}},
					Type: "",
				}},
				RefBlockBytes: "",
				RefBlockHash:  "",
				Expiration:    0,
				Timestamp:     time.Now().Unix(),
			},
			RawDataHex: "",
		},
	}

	expResp := &core.CallResult{
		TransactionList: list.ToList(),
	}

	cli.EXPECT().Call(params).Return(expResp, nil).AnyTimes()
	ch, err := obs.Observe(ctx, &w, 10)
	require.Nil(t, err)
	require.NotNil(t, ch)

	time.Sleep(time.Duration(waitInterval) * time.Millisecond)
	v, ok := <-ch
	require.Empty(t, v)
	require.Equal(t, true, ok)
	v, ok = <-ch
	require.Empty(t, v)
	require.Equal(t, false, ok)
}
