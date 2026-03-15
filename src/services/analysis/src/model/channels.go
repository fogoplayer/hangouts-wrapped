package model

// File Paths To Ingest Channel

var FilePathsToIngestChannel chan string = make(chan string)

// User Info Channel
// Skipping metrics for User Info because I'm not sure we actually use it

var UserInfoChannel chan []byte = make(chan []byte, 1)

// Chat Directory Handle Channel

var ChatDirectoryHandleChannel chan ChatDirectoryHandle = make(chan ChatDirectoryHandle)
