package core

type Wallet interface {
	GetChain() Chain
	GetPrivateKey() string
	GetAddress() string
}
