package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/jsInterface"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func main() {
	fmt.Println("Starting initialization")

	userInfoJsonChannel := make(chan jsonSchema.UserInfo_JsonSchema)
	groupInfoJsonChannel := make(chan jsonSchema.GroupInfo_JsonSchema)
	messagesJsonChannel := make(chan jsonSchema.Messages_JsonSchema)

	jsInterface.Initialize()
	fmt.Println("listening for directory")
	chatDataDirectory := <-model.ChatDataDirectoryChannel
	fmt.Println("Got directory")
	fsIo.IngestDirectory(chatDataDirectory.Name(), userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
	fmt.Println("directory channel closed")
	<-make(chan int)
}
