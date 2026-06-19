//go:build !(js && wasm)

package main

import "os"

// just args for now, can add prompting logic later if desired
func promptForChatDataDirectory() string {
	args := os.Args[1:] // exclude program
	return args[0]
}
