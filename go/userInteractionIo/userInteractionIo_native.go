//go:build !(js && wasm)

package userInteractionIo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"zarinloosli.com/hangouts-wrapped/util"
)

func prompt[T any](message string, options []string, parse func(string) (T, error)) T {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	for {
		for i, option := range options {
			fmt.Printf("  (%d) %s\n", i, option)
		}
		reader.Scan()
		text := reader.Text()
		parsed, err := parse(text)
		if err == nil {
			return parsed
		}
		fmt.Println("invalid submission, please try again")
	}
}

func Prompt(message string, options []string) int {
	return prompt(message, options, func(text string) (int, error) {
		selection, err := strconv.Atoi(text)
		if err == nil && selection >= 0 && selection < len(options) {
			return selection, nil
		}
		return -1, errors.New("")
	})
}

func MultiSelectPrompt(message string, options []string) []int {
	optionCount := len(options)

	options = append(options, "All")
	allOptionIndex := len(options) - 1
	return prompt(message, options, func(text string) ([]int, error) {
		includesAll := false
		invalidInput := false

		selectionList := strings.Split(text, ",")
		// convert list of string numbers into list of ints
		parsedSelectionList := util.ListMap(selectionList, func(selectionString string) int {
			selection, err := strconv.Atoi(selectionString)
			if err == nil && selection >= 0 && selection < len(options) {
				if selection == allOptionIndex {
					includesAll = true
				}
				return selection
			}
			invalidInput = true
			return -1
		})

		if invalidInput {
			return nil, errors.New("")
		}

		// if we encountered the "all" selection in there
		if includesAll {
			arr := []int{}
			for i := range optionCount {
				arr = append(arr, i)
			}
			return arr, nil
		}

		//otherwise
		return parsedSelectionList, nil
	})
}
