package solutions

import (
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id    int
	items []int
	op    func(item int) int
	test  func(item int) int
}

func parseToken(token string, item int) int {
	if token == "old" {
		return item
	}
	t, _ := strconv.Atoi(token)
	return t
}

func parseOp(opSpec string, mMod int) func(item int) int {
	tokens := strings.Split(opSpec, " ")
	return func(item int) int {
		a := parseToken(tokens[0], item)
		b := parseToken(tokens[2], item)
		op := tokens[1]
		if op == "+" {
			return (a + b) % mMod
		} else if op == "*" {
			return (a * b) % mMod
		}
		return ^int(0)
	}
}

func parseTest(testSpec []string) func(item int) int {
	div, _ := strconv.Atoi(strings.Split(testSpec[0], "divisible by ")[1])
	tgta, _ := strconv.Atoi(strings.Split(testSpec[1], "to monkey ")[1])
	tgtb, _ := strconv.Atoi(strings.Split(testSpec[2], "to monkey ")[1])
	return func(item int) int {
		if item%div == 0 {
			return tgta
		}
		return tgtb
	}
}

func completeRounds(monkeys *[]monkey, rounds int, div int) []int {
	inspectCount := make([]int, len(*monkeys))
	for i := 0; i < rounds; i++ {
		for mi := 0; mi < len(*monkeys); mi++ {
			for iti := 0; iti < len((*monkeys)[mi].items); iti++ {
				item := (*monkeys)[mi].op((*monkeys)[mi].items[iti])

				item = item / div
				tgt := (*monkeys)[mi].test(item)
				(*monkeys)[tgt].items = append((*monkeys)[tgt].items, item)
				inspectCount[mi] += 1
			}
			(*monkeys)[mi].items = make([]int, 0)
		}
	}
	return inspectCount
}

func parseMonkeys() []monkey {
	input := readInput("./input/day_11/p1.txt")
	monkeySpecs := strings.Split(input, "\n\n")
	monkeys := []monkey{}
	mMod := 1
	for _, monkeySpec := range monkeySpecs {
		lines := strings.Split(monkeySpec, "\n")
		div, _ := strconv.Atoi(strings.Split(lines[3], "divisible by ")[1])
		mMod *= div
	}
	for _, monkeySpec := range monkeySpecs {
		lines := strings.Split(monkeySpec, "\n")
		// id
		id, _ := strconv.Atoi(strings.Trim(strings.Split(lines[0], " ")[1], ":"))
		// items
		items := []int{}
		itemSpecs := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		for _, itemSpec := range itemSpecs {
			item, _ := strconv.Atoi(itemSpec)
			items = append(items, item)
		}
		//op
		opSpec := strings.Split(lines[2], "= ")[1]
		op := parseOp(opSpec, mMod)

		// test
		test := parseTest(lines[3:])

		// build monkey
		m := monkey{id: id, items: items, op: op, test: test}
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func Solution_11_1() int {
	monkeys := parseMonkeys()
	inspectCount := completeRounds(&monkeys, 20, 3)
	sort.Slice(inspectCount, func(i, j int) bool { return inspectCount[i] > inspectCount[j] })
	return inspectCount[0] * inspectCount[1]
}

func Solution_11_2() int {
	monkeys := parseMonkeys()
	inspectCount := completeRounds(&monkeys, 10000, 1)
	sort.Slice(inspectCount, func(i, j int) bool { return inspectCount[i] > inspectCount[j] })
	return inspectCount[0] * inspectCount[1]
}
