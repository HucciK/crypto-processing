package core

type Transaction interface {
	GetSender() string
	GetRecipient() string
	GetAmount() uint64
	GetContractAddress() string
	GetTimestamp() int64
	IsSuccessful() bool
}
