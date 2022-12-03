package solutions

var inherentPointMap = [3]int{1, 2, 3}

var resultPointMap = [3]int{0, 3, 6}

type result struct {
	a int
	b int
}

var letterMap = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"X": 0,
	"Y": 1,
	"Z": 2,
}

var winLossMap = map[result]result{
	{0, 2}: {2, 1},
	{0, 1}: {1, 0},
	{0, 0}: {0, 2},
	{1, 0}: {0, 0},
	{1, 2}: {2, 2},
	{1, 1}: {1, 1},
	{2, 1}: {1, 2},
	{2, 0}: {0, 1},
	{2, 2}: {2, 0},
}

var resultMap = map[result]int{
	{0, 2}: 0,
	{0, 1}: 2,
	{0, 0}: 1,
	{1, 0}: 0,
	{1, 2}: 2,
	{1, 1}: 1,
	{2, 1}: 0,
	{2, 0}: 2,
	{2, 2}: 1,
}

func Solution_02_1() int {
	input := readInputSlice("./input/day_02/p1.txt")
	points := 0
	for _, game := range input {
		p1d := letterMap[string(game[0])]
		p2d := letterMap[string(game[2])]
		ihVal := inherentPointMap[p2d]
		rVal := resultPointMap[resultMap[result{p1d, p2d}]]
		points += (ihVal + rVal)
	}
	return points
}

func Solution_02_2() int {
	input := readInputSlice("./input/day_02/p1.txt")
	points := 0
	for _, game := range input {
		p1d := letterMap[string(game[0])]
		wld := letterMap[string(game[2])]
		winLoss := winLossMap[result{p1d, wld}]
		ihVal := inherentPointMap[winLoss.b]
		rVal := resultPointMap[winLoss.a]
		points += (ihVal + rVal)
	}
	return points
}
