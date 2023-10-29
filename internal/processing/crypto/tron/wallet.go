package tron

import "github.com/HucciK/crypto-processing/internal/core"

type TronWallet struct {
	Chain      core.Chain
	PrivateKey string
	Address    string
	Balance    uint64
}

func (t *TronWallet) GetChain() core.Chain {
	return t.Chain
}

func (t *TronWallet) GetPrivateKey() string {
	return t.PrivateKey
}

func (t *TronWallet) GetAddress() string {
	return t.Address
}
