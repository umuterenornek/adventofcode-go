package main

import (
	"cmp"
	"fmt"
	"slices"
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
	acc := 0
	cals := []int{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			cals = append(cals, acc)
			acc = 0
			continue
		}
		cal, e := strconv.Atoi(line)
		if e != nil {
			fmt.Println("Can't convert {} to an int!", line)
		} else {
			acc += cal
		}
	}
	cals = append(cals, acc)

	slices.SortFunc(cals, func(i, j int) int {
		return cmp.Compare(j, i)
	})

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return cals[0] + cals[1] + cals[2]
	}
	// solve part 1 here

	return cals[0]
}
