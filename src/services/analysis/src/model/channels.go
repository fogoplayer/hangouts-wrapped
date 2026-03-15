package model

import (
	"fmt"
)

// File Paths To Ingest Channel

var FilePathsToIngestChannel chan string = make(chan string)

// User Info Channel
// Skipping metrics for User Info because I'm not sure we actually use it

var UserInfoChannel chan []byte = make(chan []byte, 1)

// Chat Directory Handle Channel

var ChatDirectoryHandleChannel chan ChatDirectoryHandle = make(chan ChatDirectoryHandle)

// Found - has been returned as the result of a directory call
// Parsed - has been turned into a JsonSchema object
// Ingested - has been turned into a model object
var IngestStats IngestStatsType

type IngestStatsType struct {
	FilesFound       int
	FilesParsed      int
	ChatsParsed      int
	ChatsIngested    int
	MessagesParsed   int
	MessagesIngested int
	// Found uint
	// Ingested uint
}

func (ingestStats IngestStatsType) String() string {
	return fmt.Sprint(
		"Files:", ingestStats.FilesParsed, "/", ingestStats.FilesFound, "\n",
		"Chats:", ingestStats.ChatsIngested, "/", ingestStats.ChatsParsed, "\n",
		"Messages:", ingestStats.MessagesIngested, "/", ingestStats.MessagesIngested, "\n",
	)
}
