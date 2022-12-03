package solutions

import (
	"errors"
	"fmt"
	"unicode"
)

func getPriority(c rune) (int, error) {
	if unicode.IsUpper(c) {
		return int(c) - 38, nil
	} else if unicode.IsLower(c) {
		return int(c) - 96, nil
	}
	return 0, errors.New("Invalid character")
}

func splitSack(ruckSack string) (string, string) {
	return ruckSack[:len(ruckSack)/2], ruckSack[len(ruckSack)/2:]
}

type PriorityMap map[rune]int
type PuzzleInput string
type PuzzleInputSlice []string

func getPriorityMap(s string) PriorityMap {
	set := PriorityMap{}
	for _, c := range s {
		p, err := getPriority(c)
		if err != nil {
			fmt.Printf("Could not get priority for character: %s", string(c))
		}
		set[c] = p
	}
	return set
}

func (m1 PriorityMap) union(m2 PriorityMap) PriorityMap {
	union := PriorityMap{}
	for r := range m1 {
		if p, exists := m2[r]; exists {
			union[r] = p
		}
	}
	for r, _ := range m2 {
		if p, exists := m1[r]; exists {
			union[r] = p
		}
	}
	return union
}

func (m PriorityMap) String() string {
	s := "{\n"
	for k, v := range m {
		s += fmt.Sprintf("\t%s: %d,\n", string(k), v)
	}
	s += "}"
	return s
}

func (s PuzzleInputSlice) splitGroups() ([][]string, error) {
	var chunks [][]string
	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			return nil, errors.New("Invalid input, couldn't split into groups")
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks, nil
}

func getSackOverlapPriority(s string) (int, error) {
	a, b := splitSack(s)
	aSet := getPriorityMap(a)
	bSet := getPriorityMap(b)
	union := aSet.union(bSet)
	if len(union) != 1 {
		return -1, errors.New(fmt.Sprintf("Incorrect overlap for sets: A: %s\n\nB: %s\n", aSet, bSet))
	}
	// return the single priority
	for _, p := range union {
		return p, nil
	}
	return -1, errors.New(fmt.Sprintf("Incorrect overlap for sets: A: %s\n\nB: %s\n", aSet, bSet))
}

func getGroupOverlapPriority(g []string) (int, error) {
	aSet := getPriorityMap(g[0])
	bSet := getPriorityMap(g[1])
	cSet := getPriorityMap(g[2])
	union := aSet.union(bSet)
	union = union.union(cSet)
	if len(union) != 1 {
		return -1, errors.New(fmt.Sprintf("Incorrect overlap for sets: A: %s\n\nB: %s\n", aSet, bSet))
	}
	// return the single priority
	for _, p := range union {
		return p, nil
	}
	return -1, errors.New(fmt.Sprintf("Incorrect overlap for sets: A: %s\n\nB: %s\n", aSet, bSet))
}

func Solution_03_1() int {
	input := readInputSlice("./input/day_03/p1.txt")
	sum := 0
	for _, s := range input {
		p, err := getSackOverlapPriority(s)
		if err != nil {
			fmt.Printf("Failed, couldn't get overlap for %s due to error %s", s, err)
		}
		sum += p
	}
	return sum
}

func Solution_03_2() int {
	input := readInputSlice("./input/day_03/p1.txt")
	groups, err := PuzzleInputSlice(input).splitGroups()
	if err != nil {
		fmt.Printf("Failed, couldn't get groups for %s due to error %s", input, err)
	}
	sum := 0
	for _, g := range groups {
		p, err := getGroupOverlapPriority(g)
		if err != nil {
			fmt.Printf("Failed, couldn't get overlap for %s", g)
		}
		sum += p
	}
	return sum
}
