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

func NewSimpleTransaction(from string, to string, amount int) *Transaction {
	var txInputs         []*TxInput
	var txOutputs        []*TxOutput

	txInput := &TxInput{[]byte("b4be9c2c7e771c2612263e2452495" +
		"ef3a470c747811be01eb19b0f4069795847"), 0, from}
	txInputs = append(txInputs, txInput)

	txOutput := &TxOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)
	//找零
	if amount < 10 {
		txOutput = &TxOutput{10 - amount, from}
		txOutputs = append(txOutputs, txOutput)
	}
	tx := Transaction{nil, txInputs, txOutputs}
	tx.HashTranscation()
	return &tx

}