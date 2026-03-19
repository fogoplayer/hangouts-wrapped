package main

import (
	"fmt"
	"runtime"
	"time"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/parse"
	"zarinloosli.com/hangouts-wrapped/setup"
	"zarinloosli.com/hangouts-wrapped/state"
)

func main() {
	setup.Setup()
	state.ApplicationPhase.Set(state.WaitingForDirectory)
	chatDataDirectory := promptForChatDataDirectory()
	ingestChatDirectory(chatDataDirectory)
	parseIngestedFiles()

	state.ApplicationPhase.Set(state.Ingesting)
	if runtime.GOOS != "js" {
		go func() {
			for state.ApplicationPhase.Value() == state.Ingesting {
				time.Sleep(time.Millisecond * 100)
				fmt.Println(state.GetIngestStats())
			}
		}()
	}
	state.IngestWaitGroup.Wait()
	state.ApplicationPhase.Set(state.WaitingForReport)
	fmt.Println(state.GetIngestStats())
}

// TODO is this really the right place for these functions?
func ingestChatDirectory(chatDataDirectory string) {

	fsIo.ProcessFileInWaitGoRoutine(chatDataDirectory)

	go func() { // not WaitGroup goroutines because the waitgroup is how we know to close these channels
		for pathToIngest := range state.FilePathsToIngestChannel {
			fsIo.ProcessFileInWaitGoRoutine(pathToIngest)
		}
	}()
}

func parseIngestedFiles() {
	go func() {
		for chatDirectoryHandle := range state.ChatDirectoryHandleChannel {
			parse.ParseChatDirectoryHandleInWaitGoRoutine(chatDirectoryHandle)
		}
	}()

	state.IngestWaitGroup.Go(func() {
		for userInfoBytes := range state.UserInfoChannel {
			go parse.ParseUserInfoInWaitGoRoutine(userInfoBytes)
		}
	})
}
