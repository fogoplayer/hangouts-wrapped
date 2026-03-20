package subroutines

import "zarinloosli.com/hangouts-wrapped/state"

func PostIngest() {
	close(state.FilePathsToIngestChannel)
	close(state.ChatDirectoryHandleChannel)
}
