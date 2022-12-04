package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func sumCalorieSlice(calorieSlice []string) []int {
	calorieSums := make([]int, len(calorieSlice))
	for i, inventory := range calorieSlice {
		calories := strings.Split(inventory, "\n")
		sum := 0
		for _, cals := range calories {
			if cals == "" {
				continue
			}
			calsInt, err := strconv.Atoi(cals)
			if err != nil {
				fmt.Printf("Couldn't cast calories to integer due to error: %s", err)
			}
			sum += calsInt
		}
		calorieSums[i] = sum
	}
	return calorieSums
}

func Solution_01_1() int {
	calorieInput := readInput("./input/day_01/p1.txt")
	calorieSlice := strings.Split(calorieInput, "\n\n")
	calorieSums := sumCalorieSlice(calorieSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))
	return calorieSums[0]
}

func Solution_01_2() int {
	calorieInput := readInput("./input/day_01/p1.txt")
	calorieSlice := strings.Split(calorieInput, "\n\n")
	calorieSums := sumCalorieSlice(calorieSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))
	top3Cals := 0
	for i := 0; i < 3; i++ {
		top3Cals += calorieSums[i]
	}
	return top3Cals
}
