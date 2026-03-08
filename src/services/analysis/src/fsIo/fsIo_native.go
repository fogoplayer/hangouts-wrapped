//go:build !(js && wasm)

package fsIo

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
	"zarinloosli.com/hangouts-wrapped/util"
)

var GROUPS_DIRECTORY_NAME string = "Groups"
var USER_DATA_DIRECTORY_NAME string = "Users"

func ShowDirectoryPicker(channels ...chan DirectoryHandle) chan DirectoryHandle {
	// TODO implement
	return channels[0]
}

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {

	ingestGroupData(filepath.Join(path, GROUPS_DIRECTORY_NAME))
	ingestUserData(filepath.Join(path, USER_DATA_DIRECTORY_NAME))

	return nil
}

func ingestUserData(userDataPath string) error {
	// Users -> User 123456789
	contents, err := getDirectoryContentsPaths(userDataPath)
	if err != nil {
		return err
	}

	if len(contents) != 1 {
		return errors.New("Export is not one user")
	}

	// User 1234567890 -> user_info.json
	contents, err = getDirectoryContentsPaths(userDataPath)
	if len(contents) != 1 {
		return errors.New("User directory includes more than just user info")
	}

	userInfo := contents[0]
	jsonPath := filepath.Join(userInfo, "user_info.json")
	_, err = GetJsonContents[jsonSchema.UserInfo_JsonSchema](jsonPath)
	if err != nil {
		return err
	}
	// fmt.Println(user)

	return nil
}

func ingestGroupData(groupPath string) {

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

func GetJsonContents[DestinationType any](filepath string) (DestinationType, error) {
	EMPTY := *new(DestinationType)

	userFileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return EMPTY, err
	}

	userInfo := new(DestinationType)

	err = json.Unmarshal(userFileBytes, userInfo)
	if err != nil {
		return EMPTY, err
	}

	return *userInfo, nil
}

// ////////// //
// FileHandle //
// ////////// //
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

// /////////////// //
// DirectoryHandle //
// /////////////// //
type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []model.FSAgnosticHandle {
	contents, _ := getDirectoryContentsPaths(handle.path)
	// TODO handle error
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

// //////// //
// FSHandle //
// //////// //
type FSHandle struct {
	path string
}

func (handle FSHandle) Name() string {
	return filepath.Base(handle.path)
}

func (handle FSHandle) Path() string {
	return handle.path
}

func (handle FSHandle) IsDirectory() bool {
	fileInfo, err := os.Stat(handle.path)
	if err != nil {
		panic(err)
	}

	return fileInfo.IsDir()
}

func (handle FSHandle) AsDirectoryHandle() (model.FSAgnosticDirectoryHandle, error) {
	if !handle.IsDirectory() {
		return nil, util.CreateUnableToCastFromAToBError[FSHandle, model.FSAgnosticDirectoryHandle](handle)
	}
	return DirectoryHandle{handle}, nil
}

func (handle FSHandle) AsFileHandle() (model.FSAgnosticFileHandle, error) {
	if handle.IsDirectory() {
		return nil, util.CreateUnableToCastFromAToBError[FSHandle, model.FSAgnosticFileHandle](handle)
	}
	return FileHandle{handle}, nil
}

func GetFSHandleFromPath(path string) FSHandle {
	return FSHandle{path}
}

var _ model.FSAgnosticHandle = FSHandle{} // Compile-time inheritance check
