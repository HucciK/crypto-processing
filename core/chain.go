package core

type Chain uint8

const (
	ChainNone Chain = iota
	ChainTron
)

type CallMethod uint8

const (
	CallMethodNone CallMethod = iota
	CallMethodTransactionList
)

type CallParams struct {
	Chain        Chain
	Method       CallMethod
	Wallet       Wallet
	ContractAddr string
}

type CallResult struct {
	TransactionList []Transaction
}
