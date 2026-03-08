//go:build !(js && wasm)

package fsIo

import (
	"os"

	"zarinloosli.com/hangouts-wrapped/model"
)

type FileHandle struct {
	FSHandle
}

func (handle FileHandle) Bytes() chan []byte {
	userFileBytes, _ := os.ReadFile(handle.path)
	// TODO handle error

	byteChannel := make(chan []byte, 1)
	byteChannel <- userFileBytes
	return byteChannel
}

var _ model.FSAgnosticFileHandle = FileHandle{} // Compile-time inheritance check
