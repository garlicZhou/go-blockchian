package blockchian

type TxOutput struct {
	value             int
	ScriptPubkey      string   //UTXO所有者
}