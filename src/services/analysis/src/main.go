package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/setup"
)

func main() {
	setup.Setup()

	chatDataDirectory := promptForChatDataDirectory()
	fmt.Println(chatDataDirectory)
	fsIo.ProcessFile(chatDataDirectory)

	i := 0
	for range model.BytesChannel {
		i += 1
		fmt.Println(i, "bytes recieved")
	}
}
