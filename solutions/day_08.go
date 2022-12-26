package solutions

import (
	"strconv"
	"strings"
)

func parseTrees() [][]int {
	input := readInput("./input/day_08/p1.txt")
	trees := [][]int{}
	for _, row := range strings.Split(input, "\n") {
		treeRow := []int{}
		for _, c := range row {
			treeH, _ := strconv.Atoi(string(c))
			treeRow = append(treeRow, treeH)
		}
		trees = append(trees, treeRow)
	}
	return trees
}

func countTrees(trees *[][]int) int {
	visibleCount := 0
	for i := 0; i < len((*trees)); i++ {
		for j := 0; j < len((*trees)[i]); j++ {
			visible := true
			for ie := j - 1; ie > -1; ie-- {
				if (*trees)[i][ie] >= (*trees)[i][j] {
					visible = false
					break
				}
			}
			if visible {
				visibleCount += 1
				continue
			}
			visible = true
			for ie := j + 1; ie < len((*trees)[i]); ie++ {
				if (*trees)[i][ie] >= (*trees)[i][j] {
					visible = false
					break
				}
			}
			if visible {
				visibleCount += 1
				continue
			}
			visible = true
			for ie := i - 1; ie > -1; ie-- {
				if (*trees)[ie][j] >= (*trees)[i][j] {
					visible = false
					break
				}
			}
			if visible {
				visibleCount += 1
				continue
			}
			visible = true
			for ie := i + 1; ie < len((*trees)); ie++ {
				if (*trees)[ie][j] >= (*trees)[i][j] {
					visible = false
					break
				}
			}
			if visible {
				visibleCount += 1
				continue
			}
		}
	}
	return visibleCount
}

func calculateSceneScore(trees *[][]int) [][]int {
	scores := make([][]int, len(*trees))
	for i := 0; i < len(*trees); i++ {
		scores[i] = make([]int, len((*trees)[i]))
		for j := 0; j < len((*trees)[i]); j++ {
			mult := 0
			for ie := j - 1; ie > -1; ie-- {
				mult += 1
				if (*trees)[i][ie] >= (*trees)[i][j] {
					break
				}
			}
			scores[i][j] = mult
			mult = 0
			for ie := j + 1; ie < len((*trees)[i]); ie++ {
				mult += 1
				if (*trees)[i][ie] >= (*trees)[i][j] {
					break
				}
			}
			scores[i][j] *= mult
			mult = 0
			for ie := i - 1; ie > -1; ie-- {
				mult += 1
				if (*trees)[ie][j] >= (*trees)[i][j] {
					break
				}
			}
			scores[i][j] *= mult
			mult = 0
			for ie := i + 1; ie < len((*trees)); ie++ {
				mult += 1
				if (*trees)[ie][j] >= (*trees)[i][j] {
					break
				}
			}
			scores[i][j] *= mult
		}
	}
	return scores
}

func Solution_08_1() int {
	trees := parseTrees()
	visibleCount := countTrees(&trees)
	return visibleCount
}

func Solution_08_2() int {
	trees := parseTrees()
	scores := calculateSceneScore(&trees)
	max := 0
	for _, row := range scores {
		for _, score := range row {
			if score > max {
				max = score
			}
		}
	}
	return max
}
