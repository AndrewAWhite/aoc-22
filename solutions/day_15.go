package solutions

import (
	"regexp"
	"sort"
	"strconv"
)

type sensor struct {
	x int
	y int
	d int
}

type minmax struct {
	min int
	max int
}

func parseSensor(sensorString string) sensor {
	pattern := regexp.MustCompile(`Sensor at x=([-]?\d+), y=([-]?\d+)`)
	match := pattern.FindStringSubmatch(sensorString)
	x, _ := strconv.Atoi(match[1])
	y, _ := strconv.Atoi(match[2])
	pattern = regexp.MustCompile(`beacon is at x=([-]?\d+), y=([-]?\d+)`)
	match = pattern.FindStringSubmatch(sensorString)
	bx, _ := strconv.Atoi(match[1])
	by, _ := strconv.Atoi(match[2])
	return sensor{x: x, y: y, d: abs(x-bx) + abs(y-by)}
}

func parseSensors(inputLines []string) []sensor {
	sensors := make([]sensor, len(inputLines))
	for i, sensorString := range inputLines {
		s := parseSensor(sensorString)
		sensors[i] = s
	}
	return sensors
}

func resolveFilters(filters *map[sensor]minmax, mkey sensor, fkey sensor) {
	mm, ok := (*filters)[mkey]
	if !ok {
		return
	}
	f, ok := (*filters)[fkey]
	if !ok {
		return
	}
	if mm.min <= f.min && mm.max >= f.max {
		delete((*filters), fkey)
		(*filters)[mkey] = mm
	} else if mm.max < f.min {
		return
	} else if mm.min > f.max {
		return
	} else if mm.max >= f.max {
		delete((*filters), fkey)
		mm.min = f.min
		(*filters)[mkey] = mm
	} else if mm.min <= f.min {
		delete((*filters), fkey)
		mm.max = f.max
		(*filters)[mkey] = mm
	}
}

func makeFilters(y int, sensors []sensor) map[sensor]minmax {
	filters := map[sensor]minmax{}
	for _, s := range sensors {
		dy := abs(y - s.y)
		if dy > s.d {
			continue
		}
		min := s.x - abs(s.d-dy)
		max := s.x + abs(s.d-dy)
		mm := minmax{min: min, max: max}
		if len(filters) == 0 {
			filters[s] = mm
			continue
		}
		for k, f := range filters {
			if mm.min <= f.min && mm.max >= f.max {
				delete(filters, k)
				filters[s] = mm
			} else if mm.max < f.min {
				filters[s] = mm
			} else if mm.min > f.max {
				filters[s] = mm
			} else if mm.max >= f.max {
				delete(filters, k)
				mm.min = f.min
				filters[s] = mm
			} else if mm.min <= f.min {
				delete(filters, k)
				mm.max = f.max
				filters[s] = mm
			}
		}
	}
	if len(filters) < 2 {
		return filters
	}
	keys := make([]sensor, len(filters))
	i := 0
	for key, _ := range filters {
		keys[i] = key
		i += 1
	}
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys); j++ {
			if i == j {
				continue
			}
			resolveFilters(&filters, keys[i], keys[j])
			resolveFilters(&filters, keys[j], keys[i])
		}
	}
	return filters
}

func checkRow(y int, sensors []sensor) int {
	filters := makeFilters(y, sensors)
	count := 0
	for _, mm := range filters {
		count += mm.max - mm.min
	}
	return count
}

func searchRow(y int, sensors []sensor, max int) int {
	filters := makeFilters(y, sensors)
	count := 0
	for _, mm := range filters {
		if mm.min > max {
			continue
		}
		if mm.max < 0 {
			continue
		}
		if mm.max > max {
			mm.max = max
		}
		if mm.min < 0 {
			mm.min = 0
		}
		count += mm.max - mm.min
	}
	if count > max {
		return -1
	}
	return count
}

func identifyBeacon(y int, sensors []sensor, max int) int {
	filterMap := makeFilters(y, sensors)
	filters := make([]minmax, len(filterMap))
	i := 0
	for _, f := range filterMap {
		filters[i] = f
		i += 1
	}
	sort.Slice(filters, func(i, j int) bool { return filters[i].max < filters[j].max })
	for i, f := range filters[:len(filters)-1] {
		if f.max < filters[i+1].min {
			if f.max+1 < 0 {
				continue
			} else if f.max > max {
				continue
			}
			return f.max + 1
		}
	}
	return -1
}

func Solution_15_1() int {
	inputLines := readInputSlice("./input/day_15/p1.txt")
	sensors := parseSensors(inputLines)
	return checkRow(2000000, sensors)
}

func Solution_15_2() int {
	inputLines := readInputSlice("./input/day_15/p1.txt")
	sensors := parseSensors(inputLines)
	y := 0
	max := 4000000
	for y = 0; y < max; y++ {
		blocked := searchRow(y, sensors, max)
		if blocked == max-2 {
			break
		}
	}
	x := identifyBeacon(y, sensors, max)
	return x*4000000 + y
}
