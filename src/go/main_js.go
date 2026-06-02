package main

import "zarinloosli.com/hangouts-wrapped/state"

func promptForChatDataDirectory() string {
	// TODO instantiate/re-open channels in case they were closed on a previous iteration
	return <-state.FilePathsToIngestChannel
}
