//go:build !(js && wasm)

package userInteractionIo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Prompt(message string, options []string) int {
	reader := bufio.NewReader(os.Stdin) // TODO 	bufio.NewScanner()
	fmt.Println(message)
	for {
		for i, option := range options {
			fmt.Printf("  (%d) %s\n", i, option)
		}
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")
		selection, err := strconv.Atoi(text)
		if err == nil && selection >= 0 && selection < len(options) {
			return selection
		}
		fmt.Println("invalid submission, please try again")
	}
}
