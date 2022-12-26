package solutions

import (
	"fmt"
)

type Block struct {
	i        int
	j        int
	id       string
	height   int
	up       *Block
	down     *Block
	left     *Block
	right    *Block
	distance int
}

func Hash(b Block) []byte {
	return []byte{byte(b.i), byte(b.j)}
}

var startHeightChar = "a"
var endHeightChar = "z"
var fVisited map[string]int
var bVisited map[string]int
var fQueue []*Block
var bQueue []*Block
var pathLen int

func readGrid() *[][]*Block {
	input := readInputSlice("./input/day_12/p1.txt")
	blocks := make([][]*Block, len(input))
	for i := 0; i < len(input); i++ {
		blocks[i] = make([]*Block, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			h := int(input[i][j])
			id := fmt.Sprintf("%d-%d", i, j)
			dist := int(^uint(0) >> 1)
			if string(input[i][j]) == "S" {
				h = int(startHeightChar[0])
				id = string(input[i][j])
			} else if string(input[i][j]) == "E" {
				h = int(endHeightChar[0])
				id = string(input[i][j])
			}
			blocks[i][j] = &Block{
				i:        i,
				j:        j,
				id:       id,
				height:   h,
				distance: dist,
			}
		}
	}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			block := blocks[i][j]
			up := &Block{}
			if i > 0 {
				up = blocks[i-1][j]
			}
			down := &Block{}
			if i < len(input)-1 {
				down = blocks[i+1][j]
			}
			left := &Block{}
			if j > 0 {
				left = blocks[i][j-1]
			}
			right := &Block{}
			if j < len(input[i])-1 {
				right = blocks[i][j+1]
			}
			block.up = up
			block.down = down
			block.left = left
			block.right = right
			blocks[i][j] = block
		}
	}
	return &blocks
}

func refreshGrid(blocks *[][]*Block) {
	for i := 0; i < len(*blocks); i++ {
		for j := 0; j < len((*blocks)[i]); j++ {
			dist := int(^uint(0) >> 1)
			(*blocks)[i][j].distance = dist
		}
	}
}

func findBlock(blocks [][]*Block, id string) *Block {
	start := &Block{}
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < len(blocks[i]); j++ {
			if blocks[i][j].id == id {
				start = blocks[i][j]
				break
			}
			if start.id == id {
				break
			}
		}
	}
	return start
}

func findBlocks(blocks [][]*Block, height int) []*Block {
	found := []*Block{}
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < len(blocks[i]); j++ {
			if blocks[i][j].height == height {
				found = append(found, blocks[i][j])
				break
			}
		}
	}
	return found
}

func getQueue(direction int) *[]*Block {
	if direction == 1 {
		return &fQueue
	} else if direction == -1 {
		return &bQueue
	}
	return &[]*Block{}
}

func getVisited(direction int) *map[string]int {
	if direction == 1 {
		return &fVisited
	} else if direction == -1 {
		return &bVisited
	}
	return &map[string]int{}
}

func visitNeighbour(visited *map[string]int, b *Block, prevB *Block, dist int, direction int) {
	d, ok := (*visited)[b.id]
	if ok && d <= dist {
		return
	}
	if b.id == "" {
		return
	}
	if direction == 1 && b.height > prevB.height+1 {
		return
	}
	if direction == -1 && b.height < prevB.height-1 {
		return
	}
	otherVisited := getVisited(-1 * direction)
	rev, rOk := (*otherVisited)[b.id]
	if rOk {
		if rev+dist < pathLen {
			pathLen = rev + dist
		}
		return
	}
	(*visited)[b.id] = dist
	b.distance = dist
	queue := getQueue(direction)
	*queue = append(*queue, b)
}

func visitNeighbours(b *Block, direction int, dist int) {
	visited := getVisited(direction)
	visitNeighbour(visited, b.up, b, dist, direction)
	visitNeighbour(visited, b.down, b, dist, direction)
	visitNeighbour(visited, b.left, b, dist, direction)
	visitNeighbour(visited, b.right, b, dist, direction)
}

func markDist(blocks [][]*Block, start *Block, end *Block) {
	fQueue = append(fQueue, start)
	bQueue = append(bQueue, end)
	fLen := len(fQueue)
	bLen := len(bQueue)
	for {
		fLen = len(fQueue)
		bLen = len(bQueue)
		if fLen == 0 && bLen == 0 {
			break
		}
		if fLen > 0 {
			next := popLastN(&fQueue, 1)[0]
			visitNeighbours(next, 1, next.distance+1)
		}
		if bLen > 0 {
			next := popLastN(&bQueue, 1)[0]
			visitNeighbours(next, -1, next.distance+1)
		}
	}
}

func initGlobals() {
	fQueue = []*Block{}
	bQueue = []*Block{}
	fVisited = map[string]int{}
	bVisited = map[string]int{}
}

func Solution_12_1() int {
	blocks := *readGrid()
	start := findBlock(blocks, "S")
	end := findBlock(blocks, "E")
	start.distance = 0
	end.distance = 0
	initGlobals()
	pathLen = int(^uint(0) >> 1)
	markDist(blocks, start, end)
	return pathLen
}

func Solution_12_2() int {
	blocks := *readGrid()
	lowPoints := findBlocks(blocks, int("a"[0]))
	end := findBlock(blocks, "E")
	pathLen = int(^uint(0) >> 1)
	for _, start := range lowPoints {
		refreshGrid(&blocks)
		start.distance = 0
		end.distance = 0
		markDist(blocks, start, end)
	}
	return pathLen
}
