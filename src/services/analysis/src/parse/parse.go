package parse

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/model"
)

func ParseChatDirectoryHandle(handle model.ChatDirectoryHandle) {
	fmt.Println(handle.DirectoryHandle.Name())
}

func ParseUserInfo(bytes []byte) {
	fmt.Println("Got user info")
}
