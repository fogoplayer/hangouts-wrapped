package state

import (
	"time"

	"zarinloosli.com/hangouts-wrapped/model/parsed"
)

var _ = func() struct{} {
	AllChats.OnChange(allChatsListener)
	return struct{}{}
}()

var allChatsListener = func(changed *parsed.Chat) {
	if AllChats.Includes(changed) {
		IncludedChatsFilter.Add(changed)
	} else {
		IncludedChatsFilter.Delete(changed)
	}
}

var AllChats = SetApplicationState[*parsed.Chat]{}

var MinDateFilter = ApplicationState[time.Time]{
	value: time.Time{},
}

var MaxDateFilter = ApplicationState[time.Time]{
	value: time.Now(),
}

var IncludedUsersFilter = SetApplicationState[string]{}

var IncludedChatsFilter = SetApplicationState[*parsed.Chat]{}
