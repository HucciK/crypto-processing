package payments

import (
	"github.com/HucciK/crypto-processing/core"
	"github.com/HucciK/crypto-processing/processing/config"
	"github.com/google/uuid"
	"time"
)

type PaymentsService struct {
	config config.PaymentsConfig
}

func NewPaymentsService(config config.PaymentsConfig) *PaymentsService {
	return &PaymentsService{
		config: config,
	}
}

func (s *PaymentsService) Create(recipient string, amount uint64, qr []byte) core.Payment {
	p := core.Payment{
		Id:        uuid.New(),
		Amount:    amount,
		Address:   recipient,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(time.Duration(s.config.PaymentTTL) * time.Second),
	}
	p.SetQR(qr)
	return p
}
