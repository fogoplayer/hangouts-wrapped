package main

import (
	"zarinloosli.com/hangouts-wrapped/model"
)

func promptForChatDataDirectory() string {
	// TODO instantiate/re-open channels in case they were closed on a previous iteration
	return <-model.FilePathsToIngestChannel
}
