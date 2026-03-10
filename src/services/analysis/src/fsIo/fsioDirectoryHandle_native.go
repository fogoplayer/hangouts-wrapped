//go:build !(js && wasm)

package fsIo

import (
	"os"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model"
)

type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []model.FSAgnosticHandle {
	contents, err := getDirectoryContentsPaths(handle.path)
	if err != nil {
		panic(err)
	}

	result := []model.FSAgnosticHandle{}
	for _, entry := range contents {
		result = append(result, FSHandle{entry})
	}
	return result
}

func (handle DirectoryHandle) GetEntry(name string) (model.FSAgnosticHandle, error) {
	entryPath := filepath.Join(handle.Path(), name)

	_, err := os.Stat(entryPath)
	if err != nil {
		return nil, err
	}

	return FSHandle{entryPath}, nil
}

var _ model.FSAgnosticDirectoryHandle = DirectoryHandle{} // Compile-time inheritance check
