package model

var BytesChannel chan []byte = make(chan []byte)
var FilePathsToIngestChannel chan string = make(chan string)
var ChatDirectoryHandleChannel chan ChatDirectoryHandle = make(chan ChatDirectoryHandle)
