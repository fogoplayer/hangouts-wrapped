package state

import (
	"sync"

	"zarinloosli.com/hangouts-wrapped/model"
)

var IngestWaitGroup sync.WaitGroup

var FilePathsToIngestChannel chan string = make(chan string)
var UserInfoChannel chan []byte = make(chan []byte, 1)
var ChatDirectoryHandleChannel chan model.ChatDirectoryHandle = make(chan model.ChatDirectoryHandle)
