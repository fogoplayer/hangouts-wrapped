//go:build js && wasm

package fsIo

import (
	"zarinloosli.com/hangouts-wrapped/jsInterface"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {
	jsInterface.GetDirectory()
	return nil
}
