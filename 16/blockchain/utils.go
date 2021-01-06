package blockchian

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/labstack/gommon/log"
)

func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("transfer int64 to []byte failed!  %v\n", err)
	}
	buf := buffer.Bytes()
	return buf
}

func JsonToSlice(jsonString string) []string {
	var strSlice []string
	if err := json.Unmarshal([]byte(jsonString), &strSlice); err != nil {
		log.Printf("json to []string failed %v\n", err)
	}
	return strSlice
}