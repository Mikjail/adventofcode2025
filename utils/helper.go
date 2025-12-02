package utils

import "strings"

func GetLineList(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
	}

	return lines
}
