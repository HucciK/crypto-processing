package processing

import (
	"context"
	"github.com/HucciK/crypto-processing/internal/core"
	"github.com/HucciK/crypto-processing/internal/processing/api"
	"github.com/HucciK/crypto-processing/internal/processing/config"
	"github.com/HucciK/crypto-processing/internal/processing/crypto/tron"
	"github.com/HucciK/crypto-processing/internal/processing/observer"
	"github.com/HucciK/crypto-processing/internal/processing/payments"
	"github.com/HucciK/crypto-processing/internal/processing/qr"
)

type paymentsService interface {
	Create(recipient string, amount uint64, qr []byte) core.Payment
}

type tronService interface {
	NewWallet() (core.Wallet, error)
}

type qrService interface {
	NewQR(data string) ([]byte, error)
}

type observerService interface {
	Observe(ctx context.Context, wallet core.Wallet, amount uint64) (chan struct{}, error)
}

type CryptoProcessing struct {
	payments paymentsService
	tron     tronService
	qr       qrService
	observer observerService

	config config.CryptoProcessingConfig
}

func NewCryptoProcessing(config config.CryptoProcessingConfig) *CryptoProcessing {
	cli := api.NewClient()

	payments := payments.NewPaymentsService(config.PaymentsConfig)
	tron := tron.NewTronService()
	qr := qr.NewQRService()
	observer := observer.NewObserverService(config.ObserverConfig, cli)

	return &CryptoProcessing{
		payments: payments,
		tron:     tron,
		qr:       qr,
		observer: observer,
	}
}

func (c *CryptoProcessing) CreatePayment(ctx context.Context, amount uint64) (core.Payment, chan struct{}, error) {
	wallet, err := c.tron.NewWallet()
	if err != nil {
		return core.Payment{}, nil, err
	}

	payQR, err := c.qr.NewQR(wallet.GetAddress())
	if err != nil {
		return core.Payment{}, nil, err
	}

	payment := c.payments.Create(wallet.GetAddress(), amount, payQR)

	ch, err := c.observer.Observe(ctx, wallet, amount)
	if err != nil {
		return core.Payment{}, nil, err
	}

	return payment, ch, nil
}
