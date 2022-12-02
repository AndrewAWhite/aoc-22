package solutions

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readInput() string {
	inputBytes, err := ioutil.ReadFile("./input/day_1/p1.txt")
	if err != nil {
		fmt.Printf("Couldn't read puzzle input due to error: %s", err)
	}
	inputString := string(inputBytes)
	return inputString
}

func sumCalorieSlice(calorieSlice []string) []int {
	calorieSums := make([]int, len(calorieSlice))
	for i, inventory := range calorieSlice {
		calories := strings.Split(inventory, "\n")
		sum := 0
		for _, cals := range calories  {
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

func Solution_1_1() int{
	input := readInput()
	calorieSlice := strings.Split(input, "\n\n")
	calorieSums := sumCalorieSlice(calorieSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))
	return calorieSums[0]
}

func Solution_1_2() int{
	input := readInput()
	calorieSlice := strings.Split(input, "\n\n")
	calorieSums := sumCalorieSlice(calorieSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieSums)))
	top3Cals := 0
	for i:=0; i<3; i++{
		top3Cals += calorieSums[i]
	}
	return top3Cals
}
