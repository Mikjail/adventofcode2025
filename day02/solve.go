package day02

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	ids := utils.GetInputSeparatedByComma(input)
	result := 0
	for _, id := range ids {
		id_splitted := strings.Split(id, "-")

		from_count, err := strconv.Atoi(id_splitted[0])
		if err != nil {
			continue
		}

		until_count, err := strconv.Atoi(id_splitted[1])
		if err != nil {
			continue
		}

		for i := from_count; i <= until_count; i++ {
			from_str := strconv.Itoa(i)
			digitCount := len(from_str)

			if digitCount > 0 && digitCount%2 == 0 {
				divided := len(from_str) / 2
				if from_str[divided:] == from_str[:divided] {
					if num, err := strconv.Atoi(from_str); err == nil {
						result += num
					}
				}
			}
		}
	}

	return result
}

func SolvePuzzle2(input string) int {
	ids := utils.GetInputSeparatedByComma(input)
	result := 0
	for _, id := range ids {
		id_splitted := strings.Split(id, "-")

		from_count, err := strconv.Atoi(id_splitted[0])
		if err != nil {
			continue
		}

		until_count, err := strconv.Atoi(id_splitted[1])
		if err != nil {
			continue
		}

		for i := from_count; i <= until_count; i++ {
			from_str := strconv.Itoa(i)

			if hasRepeatedPattern(from_str) {
				result += i
			}
		}
	}

	return result
}

func hasRepeatedPattern(s string) bool {
	length := len(s)

	// Try all possible pattern lengths from 1 to half the string length
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen == 0 {
			pattern := s[:patternLen]
			repetitions := length / patternLen

			// Only consider patterns that repeat at least twice
			if repetitions >= 2 {
				isRepeated := true
				for i := patternLen; i < length; i += patternLen {
					if s[i:i+patternLen] != pattern {
						isRepeated = false
						break
					}
				}
				if isRepeated {
					return true
				}
			}
		}
	}

	return false
}
