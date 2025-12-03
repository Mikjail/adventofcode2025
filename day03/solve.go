package day03

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
)

func SolvePuzzle1(input string) int {
	lines := utils.GetLineList(input)
	result := 0

	for _, line := range lines {
		count := len(line)
		max_left := 0

		for i := 0; i < count-1; i++ {
			numberLeft := string(line[i])
			for j := i + 1; j < count; j++ {
				numberRight := string(line[j])
				number, err := strconv.Atoi(numberLeft + numberRight)
				if err == nil {
					if max_left < number {
						max_left = number
					}
				}
			}
		}

		result += max_left

	}

	return result
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
