package subroutines

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/parse"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func IngestChatDirectory(chatDataDirectory string) {
	ProcessFileInWaitGoRoutine(chatDataDirectory)

	go func() { // not WaitGroup goroutines because the waitgroup is how we know to close these channels
		for pathToIngest := range state.FilePathsToIngestChannel {
			ProcessFileInWaitGoRoutine(pathToIngest)
		}
	}()
}

func ParseIngestedFiles() {
	go func() {
		for chatDirectoryHandle := range state.ChatDirectoryHandleChannel {
			parse.ParseChatDirectoryHandleInWaitGoRoutine(chatDirectoryHandle)
		}
	}()

	state.IngestWaitGroup.Go(func() {
		for userInfoBytes := range state.UserInfoChannel {
			go parse.ParseUserInfoInWaitGoRoutine(userInfoBytes)
		}
	})
}

// TODO if no match, recursively check child folders

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

func ProcessFileInWaitGoRoutine(path string) { // TODO is this the right package for this function?
	state.IngestWaitGroup.Go(func() {
		fsHandle := fsIo.GetFSHandleFromPath(path)
		if directoryHandle, err := fsHandle.AsDirectoryHandle(); err == nil {
			switch directoryHandle.Name() {
			case CHAT_DATA_DIRECTORY:
				handleDirectoryInGoRoutine(directoryHandle)
			case USERS_DIRECTORY:
				handleDirectoryInGoRoutine(directoryHandle)
			case GROUPS_DIRECTORY:
				handleDirectoryInGoRoutine(directoryHandle)
			default:
				if util.StartsWithWords(directoryHandle.Name(), DM_DIRECTORY_PREFIX, SPACE_DIRECTORY_PREFIX) {
					handleChatDirectoryInWaitGoRoutine(directoryHandle)
				}
				if util.StartsWithWords(directoryHandle.Name(), USER_PREFIX) {
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

func handleDirectoryInGoRoutine(directoryHandle model.FSAgnosticDirectoryHandle) {
	state.IngestWaitGroup.Go(func() {
		for _, entry := range directoryHandle.Entries() {
			state.IngestWaitGroup.Go(func() { // TODO is it necessary for this to be in a GoRoutine?
				state.FilePathsToIngestChannel <- entry.Path()
			})
		}
		state.IncrementStat(state.FilesParsed) // handing a directory counts as "parsing" it
	})
}

func handleChatDirectoryInWaitGoRoutine(directoryHandle model.FSAgnosticDirectoryHandle) {
	state.IngestWaitGroup.Go(func() {
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

		state.IngestWaitGroup.Go(func() {
			state.ChatDirectoryHandleChannel <- model.ChatDirectoryHandle{
				DirectoryHandle: directoryHandle,
				Messages:        messagesBytesChannel,
				GroupInfo:       groupInfoBytesChannel,
			}
		})
		state.IncrementStat(state.FilesParsed) // handing a directory counts as "parsing" it
	})
}

func handleUserInfoInWaitGoRoutine(userInfoFileHandle model.FSAgnosticFileHandle) {
	// TODO do we actually use userInfo for anything?
	state.IngestWaitGroup.Go(func() {
		state.UserInfoChannel <- <-userInfoFileHandle.Bytes()
		close(state.UserInfoChannel)
	})
}
