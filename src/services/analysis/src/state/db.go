package state

import (
	"time"

	"zarinloosli.com/hangouts-wrapped/model/parsed"
)

var AllChats = SetApplicationState[*parsed.Chat]{}

var MinDateFilter = ApplicationState[time.Time]{
	value: time.Time{},
}

var MaxDateFilter = ApplicationState[time.Time]{
	value: time.Now(),
}

var IncludedUsersFilter = SetApplicationState[string]{}

var IncludedChatsFilter = SetApplicationState[*parsed.Chat]{}
