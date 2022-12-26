package solutions

import (
	"strconv"
	"strings"
)

func register(picked *[]int, cycle int, val int, reg *int) {
	*picked = append(*picked, val*(*reg))
	*reg += 40
}

func Solution_10_1() int {
	input := readInputSlice("./input/day_10/p1.txt")
	cycle := 0
	val := 1
	picked := []int{}
	reg := 20
	for _, cmdText := range input {
		if cycle == reg {
			register(&picked, cycle, val, &reg)
		}
		cmdSplit := strings.Split(cmdText, " ")
		op := cmdSplit[0]
		if op == "noop" {
			cycle += 1
		} else {
			n, _ := strconv.Atoi(cmdSplit[1])
			cycle += 2
			if cycle >= reg {
				register(&picked, cycle, val, &reg)
			}
			val += n
		}
	}
	sum := 0
	for _, v := range picked {
		sum += v
	}
	return sum
}

func Solution_10_2() string {
	input := readInputSlice("./input/day_10/p1.txt")
	cycle := 0
	val := 1
	screen := ""
	for _, cmdText := range input {
		cmdSplit := strings.Split(cmdText, " ")
		op := cmdSplit[0]
		if op == "noop" {
			if cycle > 0 && cycle%40 == 0 {
				screen += "\n"
				cycle = 0
			}
			if val-1 <= cycle && cycle <= val+1 {
				screen += "#"
			} else {
				screen += "."
			}
			cycle += 1
		} else {
			n, _ := strconv.Atoi(cmdSplit[1])
			for i := 0; i < 2; i++ {
				if cycle > 0 && cycle%40 == 0 {
					screen += "\n"
					cycle = 0
				}
				if val-1 <= cycle && cycle <= val+1 {
					screen += "#"
				} else {
					screen += "."
				}
				cycle += 1
			}
			val += n
		}
	}
	return screen
}
