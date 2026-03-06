package setup

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/jsInterface"
)

func Setup() {
	fmt.Println("Starting initialization")
	jsInterface.Initialize()
	go (func() {})()
}
