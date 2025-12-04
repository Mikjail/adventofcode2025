package day04

import "spissable/advent-of-go-template/utils"

func SolvePuzzle1(input string) int {
	lines := utils.GetLineList(input)
	paper_roll := "@"
	arr := utils.ParseArrayStringIntoMatrix(lines)
	rows := len(arr)
	cols := len(arr[0])

	accessibleRolls := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if arr[row][col] == paper_roll {
				neighborCount := checkAdjacentPapers(arr, row, col, rows, cols, paper_roll)
				if neighborCount < 4 {
					accessibleRolls++
				}
			}
		}
	}

	return accessibleRolls
}

func checkAdjacentPapers(arr [][]string, row, col, rows, cols int, target string) int {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if arr[newRow][newCol] == target {
				count++
			}
		}
	}

	return count
}

func SolvePuzzle2(input string) int {
	lines := utils.GetLineList(input)
	paper_roll := "@"
	arr := utils.ParseArrayStringIntoMatrix(lines)
	ROWS := len(arr)
	COLS := len(arr[0])
	totalRemoved := 0

	for {
		toRemove := [][]int{}

		for row := 0; row < ROWS; row++ {
			for col := 0; col < COLS; col++ {
				if arr[row][col] == paper_roll {
					neighborCount := checkAdjacentPapers(arr, row, col, ROWS, COLS, paper_roll)
					if neighborCount < 4 {
						toRemove = append(toRemove, []int{row, col})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			arr[pos[0]][pos[1]] = "."
			totalRemoved++
		}
	}

	return totalRemoved
}
