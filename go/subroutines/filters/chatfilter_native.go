//go:build !(js && wasm)

package filters

import (
	"zarinloosli.com/hangouts-wrapped/state"
	"zarinloosli.com/hangouts-wrapped/userInteractionIo"
)

func SetChatFilter() {
	allChatNames := state.GetStableChatNamesList()

	selections := userInteractionIo.MultiSelectPrompt(
		"Enter a comma-separated list of the chats you want to include:",
		allChatNames,
	)

	SetChatFilterFromIndexes(selections)
}
