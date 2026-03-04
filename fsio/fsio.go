package fsio

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

var GROUPS_DIRECTORY_NAME string = "Groups"
var USER_DATA_DIRECTORY_NAME string = "Users"

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
	user, err := GetJsonContents[jsonSchema.UserInfo_JsonSchema](jsonPath)
	if err != nil {
		return err
	}
	fmt.Println(user)

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
