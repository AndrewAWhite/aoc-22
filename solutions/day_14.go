package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type pointPair struct {
	x int
	y int
}

type point struct {
	state int
}

func parseLines(inputLine *string) *[]*pointPair {
	pointSplit := strings.Split(*inputLine, "->")
	points := make([]*pointPair, len(pointSplit))
	for i, p := range pointSplit {
		xy := strings.Split(p, ",")
		x, _ := strconv.Atoi(strings.Trim(xy[0], " "))
		y, _ := strconv.Atoi(strings.Trim(xy[1], " "))
		points[i] = &pointPair{x: x, y: y}
	}
	return &points
}

func getEnd(mode int, dim int) func(points [][]*pointPair) int {
	return func(points [][]*pointPair) int {
		end := 0
		if mode == 0 {
			end = int(^uint(0) >> 1)
		}
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points[i]); j++ {
				if dim == 0 {
					if mode == 0 && points[i][j].x < end {
						end = points[i][j].x
					} else if mode == 1 && points[i][j].x > end {
						end = points[i][j].x
					}
				} else if dim == 1 {
					if mode == 0 && points[i][j].y < end {
						end = points[i][j].y
					} else if mode == 1 && points[i][j].y > end {
						end = points[i][j].y
					}
				}
			}
		}
		return end
	}
}

func getGrid(lines *[][]*pointPair, mode int) (*map[pointPair]*point, *pointPair, *pointPair) {
	minX := getEnd(0, 0)(*lines)
	maxX := getEnd(1, 0)(*lines)
	maxY := getEnd(1, 1)(*lines)
	min := pointPair{x: minX, y: 0}
	max := pointPair{x: maxX, y: maxY}
	points := map[pointPair]*point{}
	if mode == 1 {
		for i := min.x; i < max.x; i++ {
			points[pointPair{x: i, y: max.y + 2}] = &point{state: 1}
		}
	}
	return &points, &min, &max
}

func addRocks(lines *[][]*pointPair, grid *map[pointPair]*point) {
	for _, l := range *lines {
		for i := 0; i < len(l)-1; i++ {
			this := l[i]
			next := l[i+1]
			if this.x == next.x {
				s := this.y
				e := next.y
				if s > e {
					ne := e
					e = s
					s = ne
				}
				for j := s; j < e+1; j++ {
					(*grid)[pointPair{x: this.x, y: j}] = &point{state: 1}
				}
			} else if this.y == next.y {
				s := this.x
				e := next.x
				if s > e {
					ne := e
					e = s
					s = ne
				}
				for j := s; j < e+1; j++ {
					(*grid)[pointPair{x: j, y: this.y}] = &point{state: 1}
				}
			}
		}
	}
}

func dropSand(grid *map[pointPair]*point, min *pointPair, max *pointPair, mode int) bool {
	curPos := pointPair{x: 500, y: 0}
	nextPos := pointPair{x: 500, y: 0}
	i := 0
	for !(nextPos.x == curPos.x && nextPos.y == curPos.y && i > 0) {
		i += 1
		curPos = nextPos
		// try down
		if curPos.y == max.y && mode == 0 {
			return false
		}
		nextspot, exists := (*grid)[pointPair{x: curPos.x, y: curPos.y + 1}]
		if mode == 1 && curPos.y+1 == max.y+2 {
			nextspot, exists = &point{state: 1}, true
		}
		if !exists || nextspot.state == 0 {
			nextPos.x = curPos.x
			nextPos.y = curPos.y + 1
			continue
		}
		// try down + left
		if (curPos.x) == min.x && mode == 0 {
			return false
		}
		nextspot, exists = (*grid)[pointPair{x: curPos.x - 1, y: curPos.y + 1}]
		if mode == 1 && curPos.y+1 == max.y+2 {
			nextspot, exists = &point{state: 1}, true
		}
		if !exists || nextspot.state == 0 {
			nextPos.x = curPos.x - 1
			nextPos.y = curPos.y + 1
			// for infinite floor
			if nextPos.x < min.x && mode == 1 {
				min.x = nextPos.x
			}
			continue
		}
		// try down + right
		if (curPos.x) == max.x && mode == 0 {
			return false
		}
		nextspot, exists = (*grid)[pointPair{x: curPos.x + 1, y: curPos.y + 1}]
		if mode == 1 && curPos.y+1 == max.y+2 {
			nextspot, exists = &point{state: 1}, true
		}
		if !exists || nextspot.state == 0 {
			nextPos.x = curPos.x + 1
			nextPos.y = curPos.y + 1
			// for infinite floor
			if nextPos.x > max.x && mode == 1 {
				max.x = nextPos.x
			}
			continue
		}
	}
	(*grid)[pointPair{x: curPos.x, y: curPos.y}] = &point{state: 2}
	if curPos.x == 500 && curPos.y == 0 && mode == 1 {
		return false
	}
	return true
}

func printGrid(grid *map[pointPair]*point, min *pointPair, max *pointPair) {
	for i := min.y; i < max.y+3; i++ {
		row := make([]string, max.x)
		for j := min.x; j < max.x; j++ {
			p, exists := (*grid)[pointPair{x: j, y: i}]
			if !exists {
				row[j] = "."
			} else if p.state == 1 {
				row[j] = "#"
			} else if p.state == 2 {
				row[j] = "O"
			}
		}
		fmt.Printf("ROW %d\t", i)
		fmt.Println(strings.Join(row, ""))
	}
}

func countSand(mode int) int {
	inputLines := readInputSlice("./input/day_14/p1.txt")
	lines := make([][]*pointPair, len(inputLines))
	for i, inputLine := range inputLines {
		lines[i] = *parseLines(&inputLine)
	}
	grid, min, max := getGrid(&lines, mode)
	addRocks(&lines, grid)
	count := 0
	sandDrop := dropSand(grid, min, max, mode)
	for sandDrop {
		count += 1
		sandDrop = dropSand(grid, min, max, mode)
	}
	if mode == 1 {
		count += 1
	}
	return count
}

func Solution_14_1() int {
	return countSand(0)
}

func Solution_14_2() int {
	return countSand(1)
}
