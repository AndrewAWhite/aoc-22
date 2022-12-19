package solutions

import (
	"errors"
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
	input := readInput(filepath)
	return strings.Split(input, "\n")
}

func (s RuneSlice) SplitChunks(chunkSize int) ([]RuneSlice, error) {
	var chunks []RuneSlice
	for i := 0; i < len(s); i += chunkSize {
		end := i + 3
		if end > len(s) {
			return nil, errors.New("Invalid input, couldn't split into groups")
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks, nil
}

func Reverse(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func PrintDirTree(dir *File, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "\t"
	}
	fmt.Printf("%s - %s\t%d:\n", indent, dir.Name, dir.Total)
	for _, file := range dir.Children {
		if file.IsDir {
			continue
		}
		fmt.Printf("%s\t%d\t%s\n", indent, file.Size, file.Name)
	}
	for _, file := range dir.Children {
		if !file.IsDir {
			continue
		}
		PrintDirTree(file, depth+1)
	}
}
