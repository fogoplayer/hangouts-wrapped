package main

import (
	"fmt"

	"zarinloosli.com/hangouts-wrapped/jsInterface"
)

func main() {
	fmt.Println("Starting initialization")
	jsInterface.Initialize()
	<-make(chan int)
}
