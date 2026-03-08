//go:build !(js && wasm)

package fsIo

import (
	"os"
	"path/filepath"
)

var GROUPS_DIRECTORY_NAME string = "Groups"
var USER_DATA_DIRECTORY_NAME string = "Users"

func ShowDirectoryPicker(channels ...chan DirectoryHandle) chan DirectoryHandle {
	// TODO implement
	return channels[0]
}

func getDirectoryContentsPaths(directoryPath string) ([]string, error) {
	EMPTY := []string{}

	contents, err := os.ReadDir(directoryPath)
	if err != nil {
		return EMPTY, err
	}

	filePaths := []string{}
	for _, file := range contents {
		filePaths = append(filePaths, filepath.Join(directoryPath, file.Name()))
	}
	return filePaths, nil
}
