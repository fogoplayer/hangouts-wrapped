package fsIo

import (
	"strings"

	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/util"
)

const (
	CHAT_DATA_DIRECTORY    = "Google Chat"
	GROUPS_DIRECTORY       = "Groups"
	USERS_DIRECTORY        = "Users"
	USER_PREFIX            = "User"
	USER_INFO              = "user_info.json"
	DM_DIRECTORY_PREFIX    = "DM"
	SPACE_DIRECTORY_PREFIX = "Space"
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
			handleDirectoryInGoRoutine(directoryHandle)
		case USERS_DIRECTORY:
			handleDirectoryInGoRoutine(directoryHandle)
		case GROUPS_DIRECTORY:
			handleDirectoryInGoRoutine(directoryHandle)
		default:
			if startsWithWords(directoryHandle.Name(), DM_DIRECTORY_PREFIX, SPACE_DIRECTORY_PREFIX) {
				handleChatDirectoryInGoRoutine(directoryHandle)
			}
			if startsWithWords(directoryHandle.Name(), USER_PREFIX) {
				handleDirectoryInGoRoutine(directoryHandle)
			}
		}
	} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
		switch fileHandle.Name() {
		case USER_INFO:
			handleUserInfoInGoRoutine(fileHandle)
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

func handleDirectoryInGoRoutine(directoryHandle model.FSAgnosticDirectoryHandle) {
	go func() {
		for _, entry := range directoryHandle.Entries() {
			go func() { model.FilePathsToIngestChannel <- entry.Path() }()
		}
	}()
}

func handleChatDirectoryInGoRoutine(directoryHandle model.FSAgnosticDirectoryHandle) {
	go func() {
		messagesEntry, err := directoryHandle.GetEntry("messages.json")
		util.PanicIfError(err)
		messagesFile, err := messagesEntry.AsFileHandle()
		util.PanicIfError(err)
		messagesBytesChannel := messagesFile.Bytes()

		groupInfoEntry, err := directoryHandle.GetEntry("group_info.json")
		util.PanicIfError(err)
		groupInfoFile, err := groupInfoEntry.AsFileHandle()
		util.PanicIfError(err)
		groupInfoBytesChannel := groupInfoFile.Bytes()

		model.ChatDirectoryHandleChannel <- model.ChatDirectoryHandle{
			DirectoryHandle: directoryHandle,
			Messages:        messagesBytesChannel,
			GroupInfo:       groupInfoBytesChannel,
		}
	}()
}

func handleUserInfoInGoRoutine(fileHandle model.FSAgnosticFileHandle) {
	go func() { model.BytesChannel <- <-fileHandle.Bytes() }()
}
