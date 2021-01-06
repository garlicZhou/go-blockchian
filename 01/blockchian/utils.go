package blockchian

import (
	"bytes"
	"encoding/binary"
	"github.com/labstack/gommon/log"
)

func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("transfer int64 to []byte failed!  %v\n", err)
	}
	return nil
}