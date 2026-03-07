package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/model"
)

func promptForChatDataDirectory() string {
	chatDataDirectory := <-model.ChatDataDirectoryChannel
	fmt.Println("Goot directory")
	return chatDataDirectory.Name()

}
