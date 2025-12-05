package utils

import (
	"strings"
)

func GetLineList(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
	}

	return lines
}

func GetInputSeparatedByComma(input string) []string {
	lines := GetLineList(input)
	commaSeparated := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if len(part) > 0 {
				commaSeparated = append(commaSeparated, part)
			}
		}
	}

	return commaSeparated
}

func ParseArrayStringIntoMatrix(data []string) (result [][]string) {

	for _, val := range data {
		result = append(result, strings.Split(val, ""))
	}
	return result
}

func GetInputsByBlankSpace(input string) ([]string, []string) {
	lines := GetLineList(input)
	part1 := []string{}
	part2 := []string{}
	isPart1 := true
	for _, line := range lines {

		if len(line) == 0 {
			isPart1 = false
			continue
		}
		if isPart1 {
			part1 = append(part1, line)
		} else {
			part2 = append(part2, line)
		}
	}

	return part1, part2
}
