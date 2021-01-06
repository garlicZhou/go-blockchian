package blockchian

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/labstack/gommon/log"
)

type Transaction struct {
	TxHash       []byte
	Vins         []*TxInput
	Vouts        []*TxOutput
}

func NewCoinbaseTransaction(address string) *Transaction  {

	txInput := &TxInput{[]byte{}, -1, "system reward"}
	txOutput := &TxOutput{10,address}

	txCoinbase := &Transaction{nil,
		[]*TxInput{txInput},
		[]*TxOutput{txOutput}}
	txCoinbase.HashTranscation()

	return txCoinbase
}

func (tx *Transaction) HashTranscation()  {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(tx); err != nil {
		log.Panicf("tx hash encoded failed %v\n", err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}
