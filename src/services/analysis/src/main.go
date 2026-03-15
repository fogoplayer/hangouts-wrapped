package main

import (
	"fmt"
	"time"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/parse"
	"zarinloosli.com/hangouts-wrapped/setup"
)

func main() {
	setup.Setup()
	chatDataDirectory := promptForChatDataDirectory()
	ingestChatDirectory(chatDataDirectory)

	parseIngestedFiles()

	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(model.IngestStats)
		}
	}()

	<-make(chan struct{})
}

func ingestChatDirectory(chatDataDirectory string) {

	go fsIo.ProcessFile(chatDataDirectory)

	go func() {
		for pathToIngest := range model.FilePathsToIngestChannel {
			fsIo.ProcessFile(pathToIngest)
		}
	}()
}

// TODO inline goroutines
func parseIngestedFiles() {
	go func() {
		for chatDirectoryHandle := range model.ChatDirectoryHandleChannel {
			go parse.ParseChatDirectoryHandle(chatDirectoryHandle)
		}
	}()

	go func() {
		for userInfoBytes := range model.UserInfoChannel {
			go parse.ParseUserInfo(userInfoBytes)
		}
	}()
}
