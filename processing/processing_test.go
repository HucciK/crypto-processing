package processing

import (
	"context"
	"github.com/HucciK/crypto-processing/processing/config"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCryptoProcessing_CreatePayment(t *testing.T) {
	paymentTTL := 10
	service := NewCryptoProcessing(config.CryptoProcessingConfig{
		PaymentsConfig: config.PaymentsConfig{
			PaymentTTL: paymentTTL,
		},
		ObserverConfig: config.ObserverConfig{
			ObserveInterval: 3,
		},
	})

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(paymentTTL)*time.Second)
	payment, ch, err := service.CreatePayment(ctx, 10)
	require.Nil(t, err)
	require.NotNil(t, ch)
	require.NotEmpty(t, payment)
}
