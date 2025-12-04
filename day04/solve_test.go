package day04_test

import (
	"spissable/advent-of-go-template/day04"
	"spissable/advent-of-go-template/utils"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day04.SolvePuzzle1(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 1395)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day04.SolvePuzzle2(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 8451)
}
