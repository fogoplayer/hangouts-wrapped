package fsIo

import (
	"fmt"
	"strings"

	"zarinloosli.com/hangouts-wrapped/model"
)

const (
	CHAT_DATA_DIRECTORY    = "Google Chat"
	GROUPS_DIRECTORY       = "Groups"
	USERS_DIRECTORY        = "Users"
	USER_INFO              = "user_info.json"
	DM_DIRECTORY_PREFIX    = "DM"
	SPACE_DIRECTORY_PREFIX = "SPACE"
	GROUP_INFO             = "group_info.json"
	MESSAGES               = "messages.json"
)

func ProcessFile(
	path string,
) error {
	fsHandle := GetFSHandleFromPath(path)
	if directoryHandle, err := fsHandle.AsDirectoryHandle(); err == nil {
		switch directoryHandle.Name() {
		case CHAT_DATA_DIRECTORY:
			go handleDirectory(directoryHandle) // TODO move the goroutines inside the handlers instead?
		case USERS_DIRECTORY:
			go handleDirectory(directoryHandle)
		case GROUPS_DIRECTORY:
			go handleDirectory(directoryHandle)
		default:
			if startsWithWords(directoryHandle.Name(), DM_DIRECTORY_PREFIX, SPACE_DIRECTORY_PREFIX) {
				go handleChatDirectory(directoryHandle)
			}
		}
	} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
		switch fileHandle.Name() {
		case USER_INFO:
			fmt.Println("bytes for", fileHandle.Name())
			go func() { model.BytesChannel <- <-fileHandle.Bytes() }()
		case GROUP_INFO:
		case MESSAGES:
		default:
			// TODO probably an image file
		}
	}
	return nil
}

func startsWithWords(candidate string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		firstWord := strings.Split(candidate, " ")[0]
		if firstWord == prefix {
			return true
		}
	}
	return false
}

func handleDirectory(directoryHandle model.FSAgnosticDirectoryHandle) {
	for _, entry := range directoryHandle.Entries() {
		go func() { model.FilePathsToIngestChannel <- entry.Path() }()
	}
}

func handleChatDirectory(directoryHandle model.FSAgnosticDirectoryHandle) {
	messagesEntry, _ := directoryHandle.GetEntry("messages.json")
	messagesFile, _ := messagesEntry.AsFileHandle()
	// TODO error handling
	messagesBytesChannel := messagesFile.Bytes()

	groupInfoEntry, _ := directoryHandle.GetEntry("group_info.json")
	groupInfoFile, _ := groupInfoEntry.AsFileHandle()
	// TODO error handling
	groupInfoBytesChannel := groupInfoFile.Bytes()

	// go func() {
	// 	var messagesBytes []byte
	// 	var groupInfoBytes []byte

	// 	for range 2 {
	// 		select {
	// 		case messagesBytes = <-messagesBytesChannel:
	// 		case groupInfoBytes = <-groupInfoBytesChannel:
	// 		}
	// 	}
	model.ChatDirectoryHandleChannel <- model.ChatDirectoryHandle{
		DirectoryHandle: directoryHandle,
		Messages:        messagesBytesChannel,
		GroupInfo:       groupInfoBytesChannel,
	}

	// }()

}
