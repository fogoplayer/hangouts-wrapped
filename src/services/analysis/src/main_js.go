package main

import (
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
)

func promptForChatDataDirectory() string {
	chatDataDirectory := <-model.ChatDataDirectoryChannel
	fsIo.PathToFSHandle[chatDataDirectory.Path()] = chatDataDirectory
	return chatDataDirectory.Path()
}
