package setup

import "zarinloosli.com/hangouts-wrapped/jsInterface"

// Can't be a subroutine because the jsInterface needs to call subroutines
// so
//
// setup <- jsInterface <- subroutines
// └----------XXXXXXX----------⬏
func Setup() {
	jsInterface.Initialize()
}
