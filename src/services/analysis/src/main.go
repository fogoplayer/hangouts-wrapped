package main

import (
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/setup"
)

func main() {
	setup.Setup()

	chatDataDirectory := promptForChatDataDirectory()
	fsIo.IngestDirectory(chatDataDirectory, model.UserInfoJsonChannel, model.GroupInfoJsonChannel, model.MessagesJsonChannel)
}
