package fsIo

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func IngestDirectory(
	path string,
	userInfoJsonChannel chan<- jsonSchema.UserInfo_JsonSchema,
	groupInfoJsonChannel chan<- jsonSchema.GroupInfo_JsonSchema,
	messagesJsonChannel chan<- jsonSchema.Messages_JsonSchema,
) error {
	fmt.Println(path)
	return nil
}
