package state

import (
	"slices"

	"zarinloosli.com/hangouts-wrapped/model/parsed"
	"zarinloosli.com/hangouts-wrapped/util"
)

var _ = func() struct{} {
	AllChats.OnChange(func(c *parsed.Chat) {
		canonicalUnfilteredChatsOrder = IncludedChatsFilter.Value()
		slices.SortFunc(canonicalUnfilteredChatsOrder, func(i, j *parsed.Chat) int {
			if i.Name < j.Name {
				return -1
			} else if i.Name > j.Name {
				return 1
			} else {
				return 0
			}
		})
	})
	return struct{}{}
}()

var canonicalUnfilteredChatsOrder []*parsed.Chat

func GetFilteredChats() []*parsed.Chat {
	return IncludedChatsFilter.Value()
}

func GetStableChatNamesList() []string {
	return util.ListMap(canonicalUnfilteredChatsOrder, func(chat *parsed.Chat) string {
		return chat.Name
	})
}

func GetStableChatsList() []*parsed.Chat {
	return canonicalUnfilteredChatsOrder
}
