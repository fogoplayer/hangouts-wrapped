package main

import (
	"fmt"
	"runtime"
	"time"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/parse"
	"zarinloosli.com/hangouts-wrapped/setup"
	"zarinloosli.com/hangouts-wrapped/state"
)

func main() {
	setup.Setup()
	state.ApplicationPhase = state.WaitingForDirectory
	chatDataDirectory := promptForChatDataDirectory()
	ingestChatDirectory(chatDataDirectory)
	parseIngestedFiles()

	state.ApplicationPhase = state.Ingesting
	if runtime.GOOS != "js" {
		go func() {
			for state.ApplicationPhase == state.Ingesting {
				time.Sleep(time.Millisecond * 100)
				fmt.Println(model.GetIngestStats())
			}
		}()
	}
	model.IngestWaitGroup.Wait()
	state.ApplicationPhase = state.WaitingForReport
	fmt.Println(model.GetIngestStats())
}

// TODO is this really the right place for these functions?
func ingestChatDirectory(chatDataDirectory string) {

	fsIo.ProcessFileInWaitGoRoutine(chatDataDirectory)

	go func() { // not WaitGroup goroutines because the waitgroup is how we know to close these channels
		for pathToIngest := range model.FilePathsToIngestChannel {
			fsIo.ProcessFileInWaitGoRoutine(pathToIngest)
		}
	}()
}

func parseIngestedFiles() {
	go func() {
		for chatDirectoryHandle := range model.ChatDirectoryHandleChannel {
			parse.ParseChatDirectoryHandleInWaitGoRoutine(chatDirectoryHandle)
		}
	}()

	model.IngestWaitGroup.Go(func() {
		for userInfoBytes := range model.UserInfoChannel {
			go parse.ParseUserInfoInWaitGoRoutine(userInfoBytes)
		}
	})
}
