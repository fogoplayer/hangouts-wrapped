package fsIo

import (
	"fmt"
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

func ProcessFileInWaitGoRoutine(
	path string,
) {
	model.IngestWaitGroup.Go(func() {
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
					handleChatDirectoryInWaitGoRoutine(directoryHandle)
				}
				if startsWithWords(directoryHandle.Name(), USER_PREFIX) {
					handleDirectoryInGoRoutine(directoryHandle)
				}
			}
		} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
			switch fileHandle.Name() {
			case USER_INFO:
				handleUserInfoInWaitGoRoutine(fileHandle)
			case GROUP_INFO: // handled by ChatDirectory
			case MESSAGES: // handled by ChatDirectory
			default:
				// TODO probably an attachment file
			}
		}
	})
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
	model.IngestWaitGroup.Go(func() {
		for _, entry := range directoryHandle.Entries() {
			model.IngestWaitGroup.Go(func() { // TODO is it necessary for this to be in a GoRoutine?
				model.FilePathsToIngestChannel <- entry.Path()
			})
		}
		model.IncrementStat(model.FilesParsed) // handing a directory counts as "parsing" it
	})
}

func handleChatDirectoryInWaitGoRoutine(directoryHandle model.FSAgnosticDirectoryHandle) {
	model.IngestWaitGroup.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				// have to inline the \n to make this atomic, otherwise other goroutines will print in between
				fmt.Println("Unable to read in", directoryHandle.Path(), "\n", r)
			}
		}()

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

		model.IngestWaitGroup.Go(func() {
			model.ChatDirectoryHandleChannel <- model.ChatDirectoryHandle{
				DirectoryHandle: directoryHandle,
				Messages:        messagesBytesChannel,
				GroupInfo:       groupInfoBytesChannel,
			}
		})
		model.IncrementStat(model.FilesParsed) // handing a directory counts as "parsing" it
	})
}

func handleUserInfoInWaitGoRoutine(userInfoFileHandle model.FSAgnosticFileHandle) {
	// TODO do we actually use userInfo for anything?
	model.IngestWaitGroup.Go(func() {
		model.UserInfoChannel <- <-userInfoFileHandle.Bytes()
		close(model.UserInfoChannel)
	})
}
