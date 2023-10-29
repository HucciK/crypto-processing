package core

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	Id        uuid.UUID `json:"id"`
	Amount    uint64    `json:"amount"`
	Address   string    `json:"address"`
	qr        []byte
	IsPayed   bool      `json:"is_payed"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (p *Payment) SetQR(qr []byte) {
	p.qr = qr
}

func (p *Payment) GetQR() []byte {
	return p.qr
}
