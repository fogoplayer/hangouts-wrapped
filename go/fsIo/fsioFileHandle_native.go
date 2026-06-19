//go:build !(js && wasm)

package fsIo

import (
	"fmt"
	"os"

	"zarinloosli.com/hangouts-wrapped/model"
)

type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	userFileBytes, err := os.ReadFile(handle.path)
	if err != nil {
		panic(fmt.Errorf("Unable to read file %s", handle.Path()))
	}

	byteChannel := make(chan []byte, 1)
	byteChannel <- userFileBytes
	return byteChannel
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check
