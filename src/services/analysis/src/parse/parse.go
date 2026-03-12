package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func ParseChatDirectoryHandle(handle model.ChatDirectoryHandle) {
	groupInfoJson := jsonSchema.GroupInfo_JsonSchema{}
	messagesJson := jsonSchema.Messages_JsonSchema{}

	// TODO parallelize
	parseJson(<-handle.GroupInfo, &groupInfoJson)
	chat := parseGroupInfo(groupInfoJson)
	// fmt.Println(chat.Name)

	parseJson(<-handle.Messages, &messagesJson)

	message := "no messages"
	if len(messagesJson.Messages) > 0 {
		message = messagesJson.Messages[0].Text_
	}
	fmt.Println(chat.Name)
	fmt.Println("\t", message)
}

func ParseUserInfo(bytes []byte) {
	userInfoJson := jsonSchema.UserInfo_JsonSchema{}
	err := parseJson(bytes, &userInfoJson)
	if err != nil {
		fmt.Println("Error parsing user info:", err)
	} else {
		fmt.Println(userInfoJson.User.Name)
	}
}

func parseJson(bytes []byte, destinationPointer any) error {
	if !json.Valid(bytes) {
		return errors.New("invalid json")
	}
	json.Unmarshal(bytes, destinationPointer)
	return nil
}

func parseGroupInfo(groupInfo jsonSchema.GroupInfo_JsonSchema) model.Chat {
	chat := model.Chat{}

	for _, member := range groupInfo.Members {
		chat.Members = append(chat.Members, parseMember(member))
	}

	if groupInfo.Name != "" && groupInfo.Name != "Group Chat" {
		chat.Name = groupInfo.Name
	} else {
		memberNames := []string{}
		for _, member := range chat.Members {
			memberNames = append(memberNames, member.Name)
		}

		chat.Name = "DM with " + strings.Join(memberNames, "/")
	}

	return chat
}

func parseMember(member jsonSchema.GroupInfo_Members_JsonSchema) model.User {
	return model.User{Name: member.Name, Email: member.Email}
}
