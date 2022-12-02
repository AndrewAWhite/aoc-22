package solutions

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(filePath string) string {
	inputBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Couldn't read puzzle input due to error: %s", err)
	}
	inputString := string(inputBytes)
	return inputString
}

func readInputSlice(filepath string) []string {
	input :=  readInput(filepath)
	return strings.Split(input, "\n")
}