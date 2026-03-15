package model

import "fmt"

type IngestStatsType map[string]int
type IngestStatsKey string

const (
	FilesFound       = "FilesFound"
	FilesParsed      = "FilesParsed"
	ChatsParsed      = "ChatsParsed"
	ChatsIngested    = "ChatsIngested"
	MessagesParsed   = "MessagesParsed"
	MessagesIngested = "MessagesIngested"
)

// Found - has been returned as the result of a directory call
// Parsed - has been turned into a JsonSchema object
// Ingested - has been turned into a model object
var IngestStats = map[string]int{
	FilesFound:       0,
	FilesParsed:      0,
	ChatsParsed:      0,
	ChatsIngested:    0,
	MessagesParsed:   0,
	MessagesIngested: 0,
	// Found uint
	// Ingested uint
}

func (ingestStats IngestStatsType) String() string {
	return fmt.Sprint(
		"Files:", ingestStats[FilesFound], "/", ingestStats[FilesParsed], "\n",
		"Chats:", ingestStats[ChatsParsed], "/", ingestStats[ChatsIngested], "\n",
		"Messages:", ingestStats[MessagesParsed], "/", ingestStats[MessagesIngested], "\n",
	)
}
