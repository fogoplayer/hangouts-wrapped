//go:build !(js && wasm)

package userInteractionIo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"zarinloosli.com/hangouts-wrapped/model/reports"
	"zarinloosli.com/hangouts-wrapped/util"
)

func SelectReport() reports.ReportName { // TODO is this the right package for this function?
	keys := util.GetMapKeys(reports.ReportDescriptions)
	values := util.GetMapVals(reports.ReportDescriptions)

	selection := Prompt("Choose a report by typing a number:", values)
	if !(selection >= 0 && selection < len(keys)) {
		panic(fmt.Errorf("Prompting for report returned an invalid value: %d", selection))
	}
	selectedReport := keys[selection]
	return selectedReport
}

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
