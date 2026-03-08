package main

import (
	"zarinloosli.com/hangouts-wrapped/model"
)

func promptForChatDataDirectory() string {
	return <-model.FilePathsToIngestChannel
}
