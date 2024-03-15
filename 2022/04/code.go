package main

import (
	"strconv"
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		pair_count := 0
		for _, line := range strings.Split(input, "\n") {
			split_line := strings.Split(line, ",")
			first_pair := strings.Split(split_line[0], "-")
			second_pair := strings.Split(split_line[1], "-")
			first_pair_int := make([]int, 2)
			second_pair_int := make([]int, 2)
			first_pair_int[0], _ = strconv.Atoi(first_pair[0])
			first_pair_int[1], _ = strconv.Atoi(first_pair[1])
			second_pair_int[0], _ = strconv.Atoi(second_pair[0])
			second_pair_int[1], _ = strconv.Atoi(second_pair[1])
			if first_pair_int[0] <= second_pair_int[0] && first_pair_int[1] >= second_pair_int[0] {
				pair_count++
				continue
			}
			if first_pair_int[0] >= second_pair_int[0] && first_pair_int[0] <= second_pair_int[1] {
				pair_count++
				continue
			}
			if first_pair_int[1] >= second_pair_int[0] && first_pair_int[1] <= second_pair_int[1] {
				pair_count++
				continue
			}
			if first_pair_int[0] <= second_pair_int[1] && first_pair_int[1] >= second_pair_int[1] {
				pair_count++
			}
		}
		return pair_count
	}
	// solve part 1 here
	pair_count := 0
	for _, line := range strings.Split(input, "\n") {
		split_line := strings.Split(line, ",")
		first_pair := strings.Split(split_line[0], "-")
		second_pair := strings.Split(split_line[1], "-")
		first_pair_int := make([]int, 2)
		second_pair_int := make([]int, 2)
		first_pair_int[0], _ = strconv.Atoi(first_pair[0])
		first_pair_int[1], _ = strconv.Atoi(first_pair[1])
		second_pair_int[0], _ = strconv.Atoi(second_pair[0])
		second_pair_int[1], _ = strconv.Atoi(second_pair[1])
		if first_pair_int[0] <= second_pair_int[0] && first_pair_int[1] >= second_pair_int[1] {
			pair_count++
			continue
		}
		if first_pair_int[0] >= second_pair_int[0] && first_pair_int[1] <= second_pair_int[1] {
			pair_count++
		}
	}
	return pair_count
}
