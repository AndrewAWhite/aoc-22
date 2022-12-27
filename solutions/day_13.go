package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func splitList(listString string) []string {
	itemsStrings := []string{}
	itemString := ""
	nestCount := 0
	if listString == "" {
		return itemsStrings
	}
	for _, c := range listString[1 : len(listString)-1] {
		sc := string(c)
		if sc == "[" {
			nestCount += 1
		} else if sc == "]" {
			nestCount -= 1
		} else if sc == "," && nestCount == 0 {
			itemsStrings = append(itemsStrings, itemString)
			itemString = ""
		}
		if sc != "," || (sc == "," && nestCount != 0) {
			itemString += sc
		}
	}
	if itemString != "" {
		itemsStrings = append(itemsStrings, itemString)
	}
	return itemsStrings
}

func compare(items0 []string, items1 []string) int {
	for i, item0 := range items0 {
		if i == len(items1) {
			return 1
		}
		item1 := items1[i]
		cmp0 := item0
		cmp1 := item1
		if item0[0] != "["[0] && item1[0] != "["[0] {
			i0, _ := strconv.Atoi(item0)
			i1, _ := strconv.Atoi(item1)
			if (i0 - i1) != 0 {
				return i0 - i1
			} else {
				continue
			}
		}
		if item0[0] != "["[0] {
			cmp0 = fmt.Sprintf("[%s]", item0)
		}
		if item1[0] != "["[0] {
			cmp1 = fmt.Sprintf("[%s]", item1)
		}
		r := compare(splitList(cmp0), splitList(cmp1))
		if r != 0 {
			return r
		}
	}
	return len(items0) - len(items1)
}

func Solution_13_1() int {
	input := readInput("./input/day_13/p1.txt")
	pairs := strings.Split(input, "\n\n")
	results := make([]bool, len(pairs))
	for i, pString := range pairs {
		pair := strings.Split(pString, "\n")
		pair0 := pair[0]
		pair1 := pair[1]
		items0 := splitList(pair0)
		items1 := splitList(pair1)
		results[i] = compare(items0, items1) < 0
	}
	sum := 0
	for i, p := range results {
		if !p {
			continue
		}
		sum += i + 1
	}
	return sum
}

func Solution_13_2() int {
	input := readInput("./input/day_13/p1.txt")
	input = strings.ReplaceAll(input, "\n\n", "\n")
	inputSlice := strings.Split(input, "\n")
	inputSlice = append(inputSlice, []string{
		"[[2]]",
		"[[6]]",
	}...)
	items := make([][]string, len(inputSlice))
	for i, itemsString := range inputSlice {
		items[i] = splitList(itemsString)
	}
	sort.Slice(items, func(i, j int) bool { return compare(items[i], items[j]) < 0 })
	ans := 1
	for i, p := range items {
		if len(p) != 1 {
			continue
		}
		if p[0] == "[2]" || p[0] == "[6]" {
			ans *= (i + 1)
		}
	}
	return ans
}
