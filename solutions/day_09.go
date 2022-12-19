package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func stepTail(positions *[]int, path *map[string]bool) {
	pos := *positions
	p := *path

	xDist := pos[2] - pos[0]
	yDist := pos[3] - pos[1]
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
		pos[0] += xMov
		pos[1] += yMov
	} else if absX > 1 {
		pos[0] += xMov
	} else if absY > 1 {
		pos[1] += yMov
	}
	p[fmt.Sprintf("%d,%d", pos[0], pos[1])] = true
}

func stepHead(x int, y int, positions *[]int, path *map[string]bool) {
	pos := *positions
	if x != 0 {
		pos[2] += x
		stepTail(positions, path)
	} else if y != 0 {
		pos[3] += y
		stepTail(positions, path)
	}
}

func Solution_09_1() int {
	instructions := readInputSlice("./input/day_09/p1.txt")
	positions := []int{0, 0, 0, 0}
	path := map[string]bool{}
	path["0,0"] = true
	for _, line := range instructions {
		rSplit := strings.Split(line, " ")
		distance, _ := strconv.Atoi(rSplit[1])
		dir := rSplit[0]
		if dir == "U" {
			for i := 0; i < distance; i++ {
				stepHead(-1, 0, &positions, &path)
			}
		} else if dir == "D" {
			for i := 0; i < distance; i++ {
				stepHead(1, 0, &positions, &path)
			}
		} else if dir == "L" {
			for i := 0; i < distance; i++ {
				stepHead(0, -1, &positions, &path)
			}
		} else if dir == "R" {
			for i := 0; i < distance; i++ {
				stepHead(0, 1, &positions, &path)
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
