package utils

import "strings"

func GetLineList(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
	}

	return lines
}

func GetInputSeparatedByComma(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	commaSeparated := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
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
