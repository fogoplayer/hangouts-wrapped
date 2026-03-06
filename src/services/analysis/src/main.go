//go:build !(js && wasm)

package main

import (
	"os"
	"runtime"

	. "zarinloosli.com/hangouts-wrapped/fsIo"
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
	"zarinloosli.com/hangouts-wrapped/setup"
)

func main() {
	setup.Setup()

	userInfoJsonChannel := make(chan jsonSchema.UserInfo_JsonSchema)
	groupInfoJsonChannel := make(chan jsonSchema.GroupInfo_JsonSchema)
	messagesJsonChannel := make(chan jsonSchema.Messages_JsonSchema)

	if runtime.GOOS == "js" {
		<-messagesJsonChannel // TODO to make the program run indefinitely in WASM, remove later
	} else {
		args := os.Args[1:] // exclude program

		err := IngestDirectory(args[0], userInfoJsonChannel, groupInfoJsonChannel, messagesJsonChannel)
		if err != nil {
		}
	}
}
