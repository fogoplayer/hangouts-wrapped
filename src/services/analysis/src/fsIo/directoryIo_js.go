package fsIo

import (
	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {
	println(path)
	fileHandle := browserApis.PathToFileHandle[path]
	if fileHandle.IsDirectory() {
		directoryHandle := fileHandle.AsDirectoryHandle()
		for _, v := range directoryHandle.Entries() {
			IngestDirectory(v.RelativePath(), userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
		}
	} else {
		// TODO
	}
	return nil
}
