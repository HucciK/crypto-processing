package api

import (
	"github.com/HucciK/crypto-processing/core"
	"github.com/HucciK/crypto-processing/processing/crypto/tron"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Call(t *testing.T) {
	cli := NewClient()

	w := tron.TronWallet{
		Chain:      core.ChainTron,
		PrivateKey: "",
		Address:    "TEd6j8KG8qdZW1uESMBv233AzZXkAEPB1P",
		Balance:    0,
	}

	res, err := cli.Call(core.CallParams{
		Chain:  w.Chain,
		Wallet: &w,
		Method: core.CallMethodTransactionList,
	})
	require.Nil(t, err)
	require.NotEmpty(t, res)
}
