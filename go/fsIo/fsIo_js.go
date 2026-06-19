package fsIo

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/state"
)

var PathToFSHandle map[string]model.FSAgnosticHandle = make(map[string]model.FSAgnosticHandle)

func ShowDirectoryPicker() {
	go func() {
		jsDirectoryResult := <-browserApis.ShowDirectoryPicker()
		jsDirectoryHandle, err := jsDirectoryResult.Value()
		if err != nil {
			fmt.Println("directory picker cancelled")
			return
		}

		fsHandle := FSHandle{jsDirectoryHandle}
		dirHandle := DirectoryHandle{fsHandle}
		PathToFSHandle[jsDirectoryHandle.Path()] = dirHandle

		state.FilePathsToIngestChannel <- jsDirectoryHandle.Path()
	}()
}
