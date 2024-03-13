package main

import (
	"fmt"
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
		points := 0
		for _, line := range strings.Split(input, "\n") {
			split_line := strings.Split(line, " ")
			op_choice, my_choice := split_line[0], split_line[1]

			switch my_choice {
			case "Z":
				points += 6
				switch op_choice {
				case "A":
					points += 2
				case "B":
					points += 3
				case "C":
					points += 1
				default:
					fmt.Println("Invalid input for op_choice: {}", op_choice)
				}
			case "Y":
				points += 3
				switch op_choice {
				case "A":
					points += 1
				case "B":
					points += 2
				case "C":
					points += 3
				default:
					fmt.Println("Invalid input for op_choice: {}", op_choice)
				}
			case "X":
				points += 0
				switch op_choice {
				case "A":
					points += 3
				case "B":
					points += 1
				case "C":
					points += 2
				default:
					fmt.Println("Invalid input for op_choice: {}", op_choice)
					return -1
				}
			default:
				fmt.Println("Invalid input for my_choice: {}", my_choice)
				return -1
			}
		}

		return points
	}
	// solve part 1 here
	points := 0
	for _, line := range strings.Split(input, "\n") {
		split_line := strings.Split(line, " ")
		op_choice, my_choice := split_line[0], split_line[1]

		switch my_choice {
		case "X":
			points += 1
			switch op_choice {
			case "A":
				points += 3
			case "B":
				points += 0
			case "C":
				points += 6
			default:
				fmt.Println("Invalid input for op_choice: {}", op_choice)
			}
		case "Y":
			points += 2
			switch op_choice {
			case "A":
				points += 6
			case "B":
				points += 3
			case "C":
				points += 0
			default:
				fmt.Println("Invalid input for op_choice: {}", op_choice)
			}
		case "Z":
			points += 3
			switch op_choice {
			case "A":
				points += 0
			case "B":
				points += 6
			case "C":
				points += 3
			default:
				fmt.Println("Invalid input for op_choice: {}", op_choice)
			}
		default:
			fmt.Println("Invalid input for my_choice: {}", my_choice)
			return -1
		}

	}
	return points
}
