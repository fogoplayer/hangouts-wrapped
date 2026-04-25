//go:build !(js && wasm)

package userInteractionIo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Prompt(message string, options []string) int {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	for {
		for i, option := range options {
			fmt.Printf("  (%d) %s\n", i, option)
		}
		reader.Scan()
		text := reader.Text()
		selection, err := strconv.Atoi(text)
		if err == nil && selection >= 0 && selection < len(options) {
			return selection
		}
		fmt.Println("invalid submission, please try again")
	}
}
