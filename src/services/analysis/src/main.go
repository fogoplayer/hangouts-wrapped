package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/setup"
)

func main() {
	setup.Setup()
	go func() {
		chatDataDirectory := promptForChatDataDirectory()
		fmt.Println("chat directory", chatDataDirectory)
		fsIo.ProcessFile(chatDataDirectory)
	}()

	go func() {
		for pathToIngest := range model.FilePathsToIngestChannel {
			fsIo.ProcessFile(pathToIngest)
		}
	}()
	go func() {
		i := 0
		for chatDirectoryHandle := range model.ChatDirectoryHandleChannel {
			go func() {
				i += 1
				fmt.Println(i, "groups recieved")
				fmt.Println(string((<-chatDirectoryHandle.GroupInfo)))
			}()
		}
	}()
	<-make(chan struct{})
}
