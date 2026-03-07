package model

import (
	"zarinloosli.com/hangouts-wrapped/model/jsonSchema"
)

var UserInfoJsonChannel chan jsonSchema.UserInfo_JsonSchema = make(chan jsonSchema.UserInfo_JsonSchema)
var GroupInfoJsonChannel chan jsonSchema.GroupInfo_JsonSchema = make(chan jsonSchema.GroupInfo_JsonSchema)
var MessagesJsonChannel chan jsonSchema.Messages_JsonSchema = make(chan jsonSchema.Messages_JsonSchema)
