package tron

import (
	"encoding/hex"
	"fmt"
	"github.com/HucciK/crypto-processing/internal/core"
)

const TronTransactionStatusSuccess = "SUCCESS"

type TronTransacion struct {
	Ret        []Ret   `json:"ret"`
	TxID       string  `json:"txID"`
	RawData    RawData `json:"raw_data"`
	RawDataHex string  `json:"raw_data_hex"`
}

func (t *TronTransacion) GetSender() string {
	if len(t.RawData.Contract) == 0 {
		return ""
	}
	return t.RawData.Contract[0].Parameter.Value.OwnerAddress
}

func (t *TronTransacion) GetRecipient() string {
	if len(t.RawData.Contract) == 0 {
		return ""
	}
	return t.RawData.Contract[0].Parameter.Value.ToAddress
}

func (t *TronTransacion) IsSuccessful() bool {
	if len(t.Ret) == 0 {
		return false
	}
	return t.Ret[0].ContractRet == TronTransactionStatusSuccess
}

func (t *TronTransacion) GetAmount() uint64 {
	if len(t.RawData.Contract) == 0 {
		return 0
	}
	return t.RawData.Contract[0].Parameter.Value.Amount
}

func (t *TronTransacion) GetContractAddress() string {
	return t.RawData.Contract[0].Parameter.Value.ContractAddress
}

func (t *TronTransacion) GetTimestamp() int64 {
	return t.RawData.Timestamp
}

func (t *TronTransacion) valuesFromDataHex() error {
	if len(t.RawData.Contract) == 0 {
		return nil
	}

	if t.RawData.Contract[0].Parameter.Value.Data == "" {
		return nil
	}

	dataBytes, err := hex.DecodeString(t.RawData.Contract[0].Parameter.Value.Data)
	if err != nil {
		panic(err)
		return err
	}

	rawDataBytes := dataBytes[4:]
	toSection := rawDataBytes[:32]
	amountSection := rawDataBytes[32:]

	addrBytes := toSection[12:]
	addrBytes = append([]byte{0x41}, addrBytes...)

	t.RawData.Contract[0].Parameter.Value.ToAddress = addrFromBytes(addrBytes)
	t.RawData.Contract[0].Parameter.Value.Amount = amountFromBytes(amountSection)

	return nil
}

type Ret struct {
	ContractRet string `json:"contractRet"`
	Fee         uint64 `json:"fee"`
}

type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Timestamp     int64      `json:"timestamp"`
}

type Contract struct {
	Parameter ContractParameter `json:"parameter"`
	Type      string            `json:"type"`
}

type ContractParameter struct {
	Value   ParameterValue `json:"value"`
	TypeURL string         `json:"type_url"`
}

type ParameterValue struct {
	Data            string `json:"data"`
	Amount          uint64 `json:"amount"`
	OwnerAddress    string `json:"owner_address"`
	ToAddress       string `json:"to_address"`
	ContractAddress string `json:"contract_address"`
}

type TronTrasactions struct {
	Transactions []TronTransacion `json:"data"`
}

func (t TronTrasactions) ToList() []core.Transaction {
	var trx []core.Transaction
	for _, tronTrx := range t.Transactions {
		tt := tronTrx
		if err := tt.valuesFromDataHex(); err != nil {
			fmt.Println(err)
			continue
		}
		trx = append(trx, &tt)
	}
	return trx
}
