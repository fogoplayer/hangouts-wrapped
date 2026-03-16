package subroutines

import (
	"sync"

	"zarinloosli.com/hangouts-wrapped/state"
)

var progressPrintingWaitGroup = sync.WaitGroup{}

func PostIngest() {
	close(state.FilePathsToIngestChannel)
	close(state.ChatDirectoryHandleChannel)
	progressPrintingWaitGroup.Wait()
}
