package tron

import (
	"crypto/sha256"
	"github.com/mr-tron/base58"
	"math/big"
)

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	return h.Sum(nil)
}

func amountFromBytes(bytes []byte) uint64 {
	n := new(big.Int)
	n.SetBytes(bytes)
	return n.Uint64()
}

func addrFromBytes(bytes []byte) string {
	hashed := (s256(s256(bytes)))
	checkSum := hashed[:4]

	bytes = append(bytes, checkSum...)
	return base58.Encode(bytes)
}
