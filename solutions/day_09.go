package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func stepTail(positions *[]int, path *map[string]bool, knotOffset int, knotCount int) {
	tx := 2 * knotOffset
	ty := tx + 1
	xDist := (*positions)[2*(knotOffset-1)] - (*positions)[tx]
	yDist := (*positions)[2*(knotOffset-1)+1] - (*positions)[ty]
	absX := xDist
	xMov := 1
	if xDist < 0 {
		xMov = -1
		absX = -absX
	}
	yMov := 1
	absY := yDist
	if yDist < 0 {
		yMov = -1
		absY = -absY
	}
	if absX >= 1 && absY > 1 || absX > 1 && absY >= 1 {
		(*positions)[tx] += xMov
		(*positions)[ty] += yMov
	} else if absX > 1 {
		(*positions)[tx] += xMov
	} else if absY > 1 {
		(*positions)[ty] += yMov
	}
	if knotOffset == knotCount {
		(*path)[fmt.Sprintf("%d,%d", (*positions)[tx], (*positions)[ty])] = true
	}
}

func stepHead(x int, y int, knots int, positions *[]int, path *map[string]bool) {
	if x != 0 {
		(*positions)[0] += x
		for i := 0; i < knots; i++ {
			stepTail(positions, path, i+1, knots)
		}
	} else if y != 0 {
		(*positions)[1] += y
		for i := 0; i < knots; i++ {
			stepTail(positions, path, i+1, knots)
		}
	}
}

func doSteps(knots int) int {
	instructions := readInputSlice("./input/day_09/p1.txt")
	positions := make([]int, (1+knots)*2)
	path := map[string]bool{}
	path["0,0"] = true
	for _, line := range instructions {
		rSplit := strings.Split(line, " ")
		distance, _ := strconv.Atoi(rSplit[1])
		dir := rSplit[0]
		if dir == "U" {
			for i := 0; i < distance; i++ {
				stepHead(0, -1, knots, &positions, &path)
			}
		} else if dir == "D" {
			for i := 0; i < distance; i++ {
				stepHead(0, 1, knots, &positions, &path)
			}
		} else if dir == "L" {
			for i := 0; i < distance; i++ {
				stepHead(-1, 0, knots, &positions, &path)
			}
		} else if dir == "R" {
			for i := 0; i < distance; i++ {
				stepHead(1, 0, knots, &positions, &path)
			}
		}
	}
	count := 0
	for _, s := range path {
		if s {
			count += 1
		}
	}
	return count
}

func Solution_09_1() int {
	return doSteps(1)
}

func Solution_09_2() int {
	return doSteps(9)
}
