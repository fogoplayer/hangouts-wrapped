package fsIo

import (
	"path/filepath"

	"zarinloosli.com/hangouts-wrapped/browserApis"
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {
	fsHandle := browserApis.PathToFSHandle[path]
	if fsHandle.IsDirectory() {
		directoryHandle := fsHandle.AsDirectoryHandle()
		for _, v := range directoryHandle.Entries() {
			IngestDirectory(v.RelativePath(), userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
		}
	} else {
		// TODO
		fileHandle := fsHandle.AsFileHandle()

		if filepath.Ext(fileHandle.Name()) == ".json" {
			go func() {
				model.BytesChannel <- <-fileHandle.Data()
			}()
		}
	}
	return nil
}
