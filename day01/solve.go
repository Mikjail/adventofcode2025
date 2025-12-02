package day01

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
)

func SolvePuzzle1(input string) int {
	lines := utils.GetLineList(input)

	starting_point := 50
	counter := 0
	for _, line := range lines {
		turn := line[0]
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			continue
		}

		switch turn {
		case 'L':
			starting_point = starting_point - number
			starting_point = ((starting_point % 100) + 100) % 100

		case 'R':
			starting_point = starting_point + number
			starting_point = starting_point % 100
		}

		if starting_point == 0 {
			counter += 1
		}

	}
	return counter
}

func SolvePuzzle2(input string) int {
	lines := utils.GetLineList(input)

	starting_point := 50
	counter := 0
	for _, line := range lines {
		turn := line[0]
		number, err := strconv.Atoi(line[1:])

		if err != nil {
			continue
		}

		switch turn {
		case 'L':
			// Count how many times we pass through 0 going left
			for i := 1; i <= number; i++ {
				pos := starting_point - i
				if pos < 0 {
					pos = ((pos % 100) + 100) % 100
				}
				if pos == 0 {
					counter++
				}
			}
			starting_point = starting_point - number
			starting_point = ((starting_point % 100) + 100) % 100

		case 'R':
			for i := 1; i <= number; i++ {
				pos := (starting_point + i) % 100
				if pos == 0 {
					counter++
				}
			}
			starting_point = starting_point + number
			starting_point = starting_point % 100
		}

	}
	return counter
}
