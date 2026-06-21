package filters

import (
	"zarinloosli.com/hangouts-wrapped/model/parsed"
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/util"
)

func SetChatFilterFromIndexes(selectedChatIndexes []int) {
	selectedChatPointers := util.ListMap(selectedChatIndexes, func(chatIndex int) *parsed.Chat {
		return state.GetStableChatsList()[chatIndex]
	})

	state.IncludedChatsFilter.Overwrite(selectedChatPointers...)
}
