package solutions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RuneSlice []rune

func initStacks(stacks [][]string, l int) [][]string {
	stacks = make([][]string, l)
	for x := 0; x < l; x++ {
		stacks[x] = []string{}
	}
	return stacks
}

func getStacks(stacksDef string) ([][]string, error) {
	stacks := [][]string{}
	stacksSplit := strings.Split(stacksDef, "\n")
	for _, r := range stacksSplit[:len(stacksSplit)-1] {
		row := RuneSlice(r)
		splitRow, err := row.SplitChunks(4)
		if len(stacks) == 0 {
			stacks = initStacks(stacks, len(splitRow))
		}
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Couldn't split into stacks due to %s", err))
		}
		for j, c := range splitRow {
			if string(c[1]) == " " {
				continue
			}
			stacks[j] = append(stacks[j], string(c[1]))
		}
	}
	for i := 0; i < len(stacks); i++ {
		stacks[i] = Reverse(stacks[i])
	}
	return stacks, nil
}

func popLast(stack []string) ([]string, string) {
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, v
}

func popLastN(stack []string, n int) ([]string, []string) {
	v := stack[len(stack)-n:]
	stack = stack[:len(stack)-n]
	return stack, v
}

func parseInstructions_1(instructionsDef string, stacks [][]string) [][]string {
	instructionsSplit := strings.Split(instructionsDef, "\n")
	re, _ := regexp.Compile("move (?P<count>\\d+) from (?P<source>\\d+) to (?P<dest>\\d+)")
	for _, instruction := range instructionsSplit {
		m := re.FindStringSubmatch(instruction)
		count, _ := strconv.Atoi(m[1])
		source, _ := strconv.Atoi(m[2])
		dest, _ := strconv.Atoi(m[3])
		source -= 1
		dest -= 1
		for i := 0; i < count; i++ {
			newStack, v := popLast(stacks[source])
			stacks[source] = newStack
			stacks[dest] = append(stacks[dest], v)
		}
	}
	return stacks
}

func parseInstructions_2(instructionsDef string, stacks [][]string) [][]string {
	instructionsSplit := strings.Split(instructionsDef, "\n")
	re, _ := regexp.Compile("move (?P<count>\\d+) from (?P<source>\\d+) to (?P<dest>\\d+)")
	for _, instruction := range instructionsSplit {
		m := re.FindStringSubmatch(instruction)
		count, _ := strconv.Atoi(m[1])
		source, _ := strconv.Atoi(m[2])
		dest, _ := strconv.Atoi(m[3])
		source -= 1
		dest -= 1
		newStack, v := popLastN(stacks[source], count)
		stacks[source] = newStack
		stacks[dest] = append(stacks[dest], v...)
	}
	return stacks
}

func Solution_05_1() string {
	input := readInput("./input/day_05/p1.txt")
	inputSplit := strings.Split(input, "\n\n")
	stacks, _ := getStacks(inputSplit[0])
	stacks = parseInstructions_1(inputSplit[1], stacks)
	v := ""
	for _, s := range stacks {
		v += s[len(s)-1]
	}
	return v
}

func Solution_05_2() string {
	input := readInput("./input/day_05/p1.txt")
	inputSplit := strings.Split(input, "\n\n")
	stacks, _ := getStacks(inputSplit[0])
	stacks = parseInstructions_2(inputSplit[1], stacks)
	v := ""
	for _, s := range stacks {
		v += s[len(s)-1]
	}
	return v
}
