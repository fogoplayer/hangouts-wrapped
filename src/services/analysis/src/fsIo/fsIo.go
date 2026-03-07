package fsIo

import (
	"fmt"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model"
)

func ProcessFile(
	path string,
) error {
	fsHandle := GetFSHandleFromPath(path)
	if directoryHandle, err := fsHandle.AsDirectoryHandle(); err == nil {
		for _, v := range directoryHandle.Entries() {
			ProcessFile(v.Path())
		}
	} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
		if filepath.Ext(fileHandle.Name()) == ".json" {
			go func() {
				fmt.Println("pushing", fileHandle.Name())
				model.BytesChannel <- <-fileHandle.Bytes()
			}()
		}
	}
	return nil
}
