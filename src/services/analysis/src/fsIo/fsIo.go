package fsIo

const (
	CHAT_DATA_DIRECTORY = "Google Chat"
	GROUPS_DIRECTORY    = "Groups"
	USERS_DIRECTORY     = "Users"
	USER_INFO           = "user_info.json"
	DM_DIRECTORY        = "DM *"
	SPACE_DIRECTORY     = "SPACE *"
	GROUP_INFO          = "group_info.json"
	MESSAGES            = "messages.json"
)

func ProcessFile(
	path string,
) error {
	fsHandle := GetFSHandleFromPath(path)
	if directoryHandle, err := fsHandle.AsDirectoryHandle(); err == nil {
		switch directoryHandle.Name() {
		case CHAT_DATA_DIRECTORY:
			for _, v := range directoryHandle.Entries() {

			}
		case GROUPS_DIRECTORY:
		case USERS_DIRECTORY:
		default:
			switch "" { // TODO properly handle regex
			case DM_DIRECTORY:
			case SPACE_DIRECTORY:
			}
		}

		for _, entry := range directoryHandle.Entries() {
			ProcessFile(entry.Path())
		}
	} else if fileHandle, err := fsHandle.AsFileHandle(); err == nil {
		switch fileHandle.Name() {
		case USER_INFO:
		case GROUP_INFO:
		case MESSAGES:
		default:
			// TODO probably an image file
		}
	}
	return nil
}
