package main

import (
	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
)

func promptForChatDataDirectory() string {
	chatDataDirectory := <-model.ChatDataDirectoryChannel
	browserApis.PathToFileHandle[chatDataDirectory.RelativePath()] = &chatDataDirectory.FSHandle
	return chatDataDirectory.RelativePath()
}
