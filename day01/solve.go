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
	// TODO: solve puzzle 2
	return 0
}
