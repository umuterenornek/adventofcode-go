package main

import (
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
	// when you're ready to do part 2, remove this "not implemented" block
	stacks := [][]rune{}
	stop_parsing_stacks := false
	if part2 {
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}
			if stop_parsing_stacks {
				split_procudure := strings.Split(line, " ")
				amount, _ := strconv.Atoi(split_procudure[1])
				src, _ := strconv.Atoi(split_procudure[3])
				dst, _ := strconv.Atoi(split_procudure[5])
				src--
				dst--
				crates := make([]rune, amount)
				copy(crates, stacks[src])
				stacks[src] = stacks[src][amount:]
				stacks[dst] = append(crates, stacks[dst]...)
				continue
			}
			for i, r := range line {
				if i%4 == 1 && r != ' ' {
					// check if r is an integer
					if r >= '0' && r <= '9' {
						stop_parsing_stacks = true
						break
					}
					for len(stacks) < i/4+1 {
						stacks = append(stacks, []rune{})
					}
					stacks[i/4] = append(stacks[i/4], r)
				}
			}
		}
		msg := ""
		for _, stack := range stacks {
			msg += string(stack[0])
		}
		return msg
	}
	// solve part 1 here
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if stop_parsing_stacks {
			split_procudure := strings.Split(line, " ")
			amount, _ := strconv.Atoi(split_procudure[1])
			src, _ := strconv.Atoi(split_procudure[3])
			dst, _ := strconv.Atoi(split_procudure[5])
			src--
			dst--
			crates := make([]rune, amount)
			copy(crates, stacks[src])
			slices.Reverse(crates)
			stacks[src] = stacks[src][amount:]
			stacks[dst] = append(crates, stacks[dst]...)
			continue
		}
		for i, r := range line {
			if i%4 == 1 && r != ' ' {
				// check if r is an integer
				if r >= '0' && r <= '9' {
					stop_parsing_stacks = true
					break
				}
				for len(stacks) < i/4+1 {
					stacks = append(stacks, []rune{})
				}
				stacks[i/4] = append(stacks[i/4], r)
			}
		}
	}
	msg := ""
	for _, stack := range stacks {
		msg += string(stack[0])
	}
	return msg
}
