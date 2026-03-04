package main

import (
	"os"

	. "zarinloosli.com/hangouts-wrapped/fsio"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

func main() {
	args := os.Args[1:] // exclude program

	userInfoJsonChannel := make(chan jsonSchema.UserInfo_JsonSchema)
	groupInfoJsonChannel := make(chan jsonSchema.GroupInfo_JsonSchema)
	messagesJsonChannel := make(chan jsonSchema.Messages_JsonSchema)

	err := IngestDirectory(args[0], userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
	if err != nil {
	}
}
