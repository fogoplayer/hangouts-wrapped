package browserApis

import (
	"encoding/binary"
	"errors"
	"syscall/js"
)

func GetIntFromJsValue(jsValue js.Value) (int64, error) {
	valueBytes := []byte{}
	js.CopyBytesToGo(valueBytes, jsValue)

	value, bytesRead := binary.Varint(valueBytes)
	if bytesRead <= 0 {
		return 0, errors.New("Couldn't read int")
	}
	return value, nil
}
