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
	lines := utils.GetLineList(input)
	result := 0

	for _, line := range lines {
		maxJoltage := findMaxWith12HighestDigits(line)
		result += maxJoltage
	}

	return result
}

func findMaxWith12HighestDigits(s string) int {
	digits := []byte{}

	// Extract all digit characters
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		}
	}

	if len(digits) == 0 {
		return 0
	}

	if len(digits) <= 12 {
		result, err := strconv.Atoi(string(digits))
		if err != nil {
			return 0
		}
		return result
	}

	stack := []byte{}
	toRemove := len(digits) - 12

	for _, digit := range digits {
		for len(stack) > 0 && stack[len(stack)-1] < digit && toRemove > 0 {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}

	// If we still need to remove digits, remove from the end
	for toRemove > 0 && len(stack) > 12 {
		stack = stack[:len(stack)-1]
		toRemove--
	}

	// Take only the first 12 digits
	if len(stack) > 12 {
		stack = stack[:12]
	}

	result, err := strconv.Atoi(string(stack))
	if err != nil {
		return 0
	}

	return result
}
