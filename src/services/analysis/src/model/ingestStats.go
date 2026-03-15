package model

import (
	"fmt"
	"sync"
)

type IngestStatsType map[IngestStatsKey]int
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
var ingestStats = map[IngestStatsKey]int{
	FilesFound:       0,
	FilesParsed:      0,
	ChatsParsed:      0,
	ChatsIngested:    0,
	MessagesParsed:   0,
	MessagesIngested: 0,
}

func (ingestStats IngestStatsType) String() string {
	return fmt.Sprint(
		"Files:", ingestStats[FilesFound], "/", ingestStats[FilesParsed], "\n",
		"Chats:", ingestStats[ChatsParsed], "/", ingestStats[ChatsIngested], "\n",
		"Messages:", ingestStats[MessagesParsed], "/", ingestStats[MessagesIngested], "\n",
	)
}

var ingestStatsMutex = sync.RWMutex{}

func IncrementStat(ingestStatsKey IngestStatsKey) {
	ingestStatsMutex.Lock()
	defer func() { ingestStatsMutex.Unlock() }()
	ingestStats[ingestStatsKey] += 1
}

func GetIngestStats() IngestStatsType {
	ingestStatsMutex.RLock()
	defer func() { ingestStatsMutex.RUnlock() }()
	return ingestStats
}
