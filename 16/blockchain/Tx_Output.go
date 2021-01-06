package blockchian

type TxOutput struct {
	Value             int
	ScriptPubkey      string   //UTXO所有者
}