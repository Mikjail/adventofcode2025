package day06

import (
	"spissable/advent-of-go-template/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {

	lines := utils.GetLineList(input)

	matrix := utils.SplitLinesWithNumbersIntoMatrix(lines[:len(lines)-1])

	operations := strings.Fields(lines[len(lines)-1])

	ROWS := len(matrix)
	COLS := len(matrix[0])
	result := 0
	for col := range COLS {
		operation := operations[col]
		accum, _ := strconv.Atoi(matrix[0][col])
		for j := 1; j < ROWS; j++ {
			numberB, _ := strconv.Atoi(matrix[j][col])
			accum = operateIn(operation, accum, numberB)
		}
		result += accum
	}

	return result
}

func operateIn(operation string, a, b int) int {
	switch operation {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}

}

func SolvePuzzle2(input string) int {
	lines := utils.GetLineList(input)
	matrix := utils.ParseArrayStringIntoMatrix(lines[:len(lines)-1])
	operations := strings.Fields(lines[len(lines)-1])

	rows := len(matrix)
	cols := len(matrix[0]) - 1
	results := []int{}
	accumulated := 0
	operationIdx := len(operations) - 1

	for col := cols; col >= 0; col-- {
		columnValue := getColumnValue(matrix, col, rows)

		if columnValue == "" {
			results = append(results, accumulated)
			accumulated = 0
			operationIdx--
			continue
		}

		number, err := strconv.Atoi(columnValue)
		if err != nil {
			continue
		}

		if accumulated == 0 {
			accumulated = number
		} else {
			accumulated = operateIn(operations[operationIdx], accumulated, number)
		}

		if col == 0 {
			results = append(results, accumulated)
		}
	}

	// Sum all problem results
	totalResult := 0
	for _, result := range results {
		totalResult += result
	}

	return totalResult
}

// getColumnValue combines all non-space values in a column into a single number string
func getColumnValue(matrix [][]string, col int, rows int) string {
	columnStr := ""
	for row := 0; row < rows; row++ {
		cell := matrix[row][col]
		if cell != " " {
			columnStr = strings.TrimSpace(columnStr + cell)
		}
	}
	return columnStr
}
