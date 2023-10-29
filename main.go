package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mr-tron/base58"
	"math/big"
)

func main() {
	//ctrct := "4142a1e39aefa49290f2b3f9ed688d7cecf86cd6e0"

	data := "a9059cbb000000000000000000000000076abaa36a89610946e77273a704e2231c3897200000000000000000000000000000000000000000000000000000000000a7d8c0"

	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}

	rawDataBytes := dataBytes[4:]
	toSection := rawDataBytes[:32]
	amountSection := rawDataBytes[32:]

	addrBytes := toSection[12:]
	addrBytes = append([]byte{0x41}, addrBytes...)

	toAddr := addrFromBytes(addrBytes)
	amount := amountFromBytes(amountSection)
	fmt.Println("to addr: ", toAddr, "Amount", amount)
}
func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	return h.Sum(nil)
}

func amountFromBytes(bytes []byte) string {
	n := new(big.Int)
	n.SetBytes(bytes)
	return n.String()
}

func addrFromBytes(bytes []byte) string {
	hashed := (s256(s256(bytes)))
	checkSum := hashed[:4]

	for _, c := range checkSum {
		bytes = append(bytes, c)
	}
	return base58.Encode(bytes)
}
