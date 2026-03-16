package state

// TODO create model/stats subpackage

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
	ingestStatsMutex.RLock()
	defer func() { ingestStatsMutex.RUnlock() }()
	return fmt.Sprint(
		"Files:", ingestStats[FilesParsed], "/", ingestStats[FilesFound], "\n",
		"Chats:", ingestStats[ChatsIngested], "/", ingestStats[ChatsParsed], "\n",
		"Messages:", ingestStats[MessagesIngested], "/", ingestStats[MessagesParsed], "\n",
	)
}

var ingestStatsMutex = sync.RWMutex{}

func IncrementStat(key IngestStatsKey, amounts ...int) {
	ingestStatsMutex.Lock()
	defer func() { ingestStatsMutex.Unlock() }()
	if len(amounts) >= 1 {
		ingestStats[key] += amounts[0]
	} else {
		ingestStats[key] += 1
	}
}

func GetIngestStats() IngestStatsType {
	ingestStatsMutex.RLock()
	defer func() { ingestStatsMutex.RUnlock() }()
	return ingestStats
}
