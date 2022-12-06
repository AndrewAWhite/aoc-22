package solutions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RuneSlice []rune

func initStacks(_stack *[][]string, l int) {
	stacks := *_stack
	stacks = make([][]string, l)
	for x := 0; x < l; x++ {
		stacks[x] = []string{}
	}
	*_stack = stacks
}

func parseStacks(stacksDef string) ([][]string, error) {
	stacks := [][]string{}
	stacksSplit := strings.Split(stacksDef, "\n")
	for _, r := range stacksSplit[:len(stacksSplit)-1] {
		row := RuneSlice(r)
		splitRow, err := row.SplitChunks(4)
		if len(stacks) == 0 {
			initStacks(&stacks, len(splitRow))
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

func popLastN(_stack *[]string, n int) []string {
	stack := *_stack
	v := stack[len(stack)-n:]
	*_stack = stack[:len(stack)-n]
	return v
}

func execInstructions(instructionsDef string, _stacks *[][]string, q int) {
	stacks := *_stacks
	instructionsSplit := strings.Split(instructionsDef, "\n")
	re, _ := regexp.Compile("move (?P<count>\\d+) from (?P<source>\\d+) to (?P<dest>\\d+)")
	for _, instruction := range instructionsSplit {
		m := re.FindStringSubmatch(instruction)
		count, _ := strconv.Atoi(m[1])
		source, _ := strconv.Atoi(m[2])
		dest, _ := strconv.Atoi(m[3])
		source -= 1
		dest -= 1
		if q == 1 {
			for i := 0; i < count; i++ {
				v := popLastN(&stacks[source], 1)
				stacks[dest] = append(stacks[dest], v...)
			}
		} else if q == 2 {
			v := popLastN(&stacks[source], count)
			stacks[dest] = append(stacks[dest], v...)
		}
	}
}

func solution(q int) string {
	input := readInput("./input/day_05/p1.txt")
	inputSplit := strings.Split(input, "\n\n")
	stacks, _ := parseStacks(inputSplit[0])
	execInstructions(inputSplit[1], &stacks, q)
	v := ""
	for _, s := range stacks {
		v += s[len(s)-1]
	}
	return v
}

func Solution_05_1() string {
	return solution(1)
}

func Solution_05_2() string {
	return solution(2)
}
