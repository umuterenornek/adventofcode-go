package main

import (
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
		return "not implemented"
	}
	// solve part 1 here
	buf := make([]rune, 4)
	buf_set := map[rune]struct{}{}
	for i, r := range input {
		buf[i%4] = r
		if i > 3 {
			for _, r := range buf {
				buf_set[r] = struct{}{}
			}
			if len(buf_set) == 4 {
				return i + 1
			}
			buf_set = map[rune]struct{}{}
		}
	}
	return buf
}
