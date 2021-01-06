package blockchian

type TxInput struct {
	TxHash        []byte
	Vout          int      //上一笔交易输出索引
	ScriptSig     string   //发送者用户
}
