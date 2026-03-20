package subroutines

import (
	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/parse"
	"zarinloosli.com/hangouts-wrapped/state"
)

func IngestChatDirectory(chatDataDirectory string) {
	fsIo.ProcessFileInWaitGoRoutine(chatDataDirectory)

	go func() { // not WaitGroup goroutines because the waitgroup is how we know to close these channels
		for pathToIngest := range state.FilePathsToIngestChannel {
			fsIo.ProcessFileInWaitGoRoutine(pathToIngest)
		}
	}()
}

func ParseIngestedFiles() {
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

func PostIngest() {
	close(state.FilePathsToIngestChannel)
	close(state.ChatDirectoryHandleChannel)
}
