package model

var FilePathsToIngestChannel chan string = make(chan string)
var UserInfoChannel chan []byte = make(chan []byte, 1)
var ChatDirectoryHandleChannel chan ChatDirectoryHandle = make(chan ChatDirectoryHandle)
