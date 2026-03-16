package main

import (
	"fmt"
	"runtime"
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

	if runtime.GOOS != "js" {
		stillIngesting := true
		go func() {
			for stillIngesting {
				time.Sleep(time.Millisecond * 100)
				fmt.Println(model.GetIngestStats())
			}
		}()
		model.IngestWaitGroup.Wait()
		stillIngesting = false
		fmt.Println("done looping")
	}

	model.IngestWaitGroup.Wait()
	fmt.Println("done ingesting")
}

// TODO is this really the right place for these functions?
func ingestChatDirectory(chatDataDirectory string) {

	fsIo.ProcessFileInWaitGoRoutine(chatDataDirectory)

	model.IngestWaitGroup.Go(func() {
		for pathToIngest := range model.FilePathsToIngestChannel {
			fsIo.ProcessFileInWaitGoRoutine(pathToIngest)
		}
	})
}

func parseIngestedFiles() {
	model.IngestWaitGroup.Go(func() {
		for chatDirectoryHandle := range model.ChatDirectoryHandleChannel {
			parse.ParseChatDirectoryHandleInWaitGoRoutine(chatDirectoryHandle)
		}
	})

	model.IngestWaitGroup.Go(func() {
		for userInfoBytes := range model.UserInfoChannel {
			go parse.ParseUserInfoInWaitGoRoutine(userInfoBytes)
		}
	})
}
