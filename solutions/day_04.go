package solutions

import (
	"strconv"
	"strings"
)

type pair struct {
	e1 [2]int
	e2 [2]int
}

func splitPairs(input []string) []pair {
	pairs := make([]pair, len(input))
	for i, l := range input {
		lSplit := strings.Split(l, ",")
		p1Split := strings.Split(lSplit[0], "-")
		p2Split := strings.Split(lSplit[1], "-")
		p10, _ := strconv.Atoi(p1Split[0])
		p11, _ := strconv.Atoi(p1Split[1])
		p20, _ := strconv.Atoi(p2Split[0])
		p21, _ := strconv.Atoi(p2Split[1])
		pairs[i] = pair{[2]int{p10, p11}, [2]int{p20, p21}}
	}
	return pairs
}

func checkPairContained(p pair) bool {
	// Check if e1 is fully contained in e2
	if p.e1[0] >= p.e2[0] && p.e1[1] <= p.e2[1] {
		return true
	} else if p.e2[0] >= p.e1[0] && p.e2[1] <= p.e1[1] {
		// Check the inverse
		return true
	}
	return false
}

func checkPairOverlap(p pair) bool {
	if p.e2[0] <= p.e1[1] && p.e1[0] <= p.e2[0] {
		return true
	}
	if p.e1[0] <= p.e2[1] && p.e2[0] <= p.e1[0] {
		return true
	}
	return false
}

func Solution_04_1() int {
	input := readInputSlice("./input/day_04/p1.txt")
	pairs := splitPairs(input)
	c := 0
	for _, p := range pairs {
		contained := checkPairContained(p)
		if !contained {
			continue
		}
		c += 1
	}
	return c
}

func Solution_04_2() int {
	input := readInputSlice("./input/day_04/p1.txt")
	pairs := splitPairs(input)
	c := 0
	for _, p := range pairs {
		contained := checkPairOverlap(p)
		if !contained {
			continue
		}
		c += 1
	}
	return c
}
