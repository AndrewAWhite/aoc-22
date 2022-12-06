package solutions

func findMarker(input string, n int) int {
	for i := 0; i < len(input); i++ {
		candidate := input[i : i+n]
		h := map[rune]int{}
		for _, char := range candidate {
			count := h[char]
			h[char] = count + 1
		}
		p := true
		for _, count := range h {
			if count > 1 {
				p = false
				break
			}
		}
		if p {
			return i + n
		}
	}
	return -1
}

func Solution_06_1() int {
	input := readInput("./input/day_06/p1.txt")
	v := findMarker(input, 4)
	return v
}

func Solution_06_2() int {
	input := readInput("./input/day_06/p1.txt")
	v := findMarker(input, 14)
	return v
}
