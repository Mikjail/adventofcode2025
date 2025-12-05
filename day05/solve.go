package day05

import (
	"sort"
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	part1, part2 := utils.GetInputsByBlankSpace(input)
	isFresh := 0
	for _, id := range part2 {
		number, err := strconv.Atoi(id)
		if err != nil {
			break
		}
		for _, ranges := range part1 {
			from, to := getRanges(ranges)
			if from <= number && number <= to {
				isFresh++
				break
			}
		}

	}

	return isFresh
}

func getRanges(ranges string) (int, int) {
	ranges_splitted := strings.Split(ranges, "-")
	from, err1 := strconv.Atoi(ranges_splitted[0])
	to, err2 := strconv.Atoi(ranges_splitted[1])
	if err1 != nil || err2 != nil {
		return 0, 0
	}

	return from, to
}

func SolvePuzzle2(input string) int {
	part1, _ := utils.GetInputsByBlankSpace(input)

	var ranges [][2]int
	for _, r := range part1 {
		from, to := getRanges(r)
		ranges = append(ranges, [2]int{from, to})
	}
	// Sort all ranges, so I don't have to repeat counts
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	result := 0
	currentStart, currentEnd := ranges[0][0], ranges[0][1]

	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= currentEnd+1 {
			if ranges[i][1] > currentEnd {
				currentEnd = ranges[i][1]
			}
		} else {
			// Gap found, add previous range length
			result += currentEnd - currentStart + 1
			currentStart, currentEnd = ranges[i][0], ranges[i][1]
		}
	}

	result += currentEnd - currentStart + 1

	return result
}
