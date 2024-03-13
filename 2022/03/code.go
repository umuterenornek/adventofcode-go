package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	getPriority := func(c rune) int {
		if int(c) < 97 {
			return int(c) - 38
		}
		return int(c) - 96
	}
	stringToSet := func(s string) map[rune]struct{} {
		set := make(map[rune]struct{})
		for _, char := range s {
			set[char] = struct{}{}
		}
		return set
	}
	intersection := func(set1, set2 map[rune]struct{}) map[rune]struct{} {
		intersection := make(map[rune]struct{})
		for value := range set1 {
			if _, exists := set2[value]; exists {
				intersection[value] = struct{}{}
			}
		}
		return intersection
	}
	getOnlyKey := func(m map[rune]struct{}) (key rune) {
		for k := range m {
			key = k
			break
		}
		return key
	}
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sum := 0
		rucksack := [3]map[rune]struct{}{}
		for i, line := range strings.Split(input, "\n") {
			rucksack[i%3] = stringToSet(line)
			if (i+1)%3 == 0 {
				sum += getPriority(getOnlyKey(intersection(intersection(rucksack[0], rucksack[1]), rucksack[2])))
				rucksack = [3]map[rune]struct{}{}
			}
		}
		return sum
	}
	// solve part 1 here
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		mid_idx := len(line)
		first_half_set := stringToSet(line[:mid_idx/2])
		second_half_set := stringToSet(line[mid_idx/2:])
		sum += getPriority(getOnlyKey(intersection(first_half_set, second_half_set)))
	}
	return sum
}
