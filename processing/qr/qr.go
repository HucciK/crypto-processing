package qr

import (
	"github.com/skip2/go-qrcode"
)

type QRService struct {
}

func NewQRService() *QRService {
	return &QRService{}
}

func (s *QRService) NewQR(data string) ([]byte, error) {
	return qrcode.Encode(data, qrcode.Medium, 256)
}
