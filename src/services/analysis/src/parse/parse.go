package parse

import (
	"encoding/json"
	"errors"
	"fmt"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func ParseChatDirectoryHandle(handle model.ChatDirectoryHandle) {
	// fmt.Println(handle.DirectoryHandle.Name())
}

func ParseUserInfo(bytes []byte) {
	userInfoJson := jsonSchema.UserInfo_JsonSchema{}
	err := parseJson(bytes, &userInfoJson)
	if err != nil {
		fmt.Println("Error parsing user info:", err)
	} else {
		fmt.Println(userInfoJson)
	}
}

func parseJson(bytes []byte, destinationPointer any) error {
	if !json.Valid(bytes) {
		return errors.New("invalid json")
	}
	json.Unmarshal(bytes, destinationPointer)
	return nil
}
