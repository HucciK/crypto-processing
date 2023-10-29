package tron

import (
	"encoding/hex"
	"github.com/HucciK/crypto-processing/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

type TronService struct {
}

func NewTronService() *TronService {
	return &TronService{}
}

func (s *TronService) NewWallet() (core.Wallet, error) {
	var w TronWallet
	w.Chain = core.ChainTron

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return &w, err
	}
	w.PrivateKey = hex.EncodeToString(crypto.FromECDSA(privateKey))

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	address = "41" + address[2:]

	decodeAddr, err := hex.DecodeString(address)
	if err != nil {
		return &w, err
	}

	hash1 := s256(s256(decodeAddr))
	secret := hash1[:4]

	decodeAddr = append(decodeAddr, secret...)
	w.Address = base58.Encode(decodeAddr)

	return &w, nil
}
