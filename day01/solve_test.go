package day01_test

import (
	"spissable/advent-of-go-template/day01"
	"spissable/advent-of-go-template/utils"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day01.SolvePuzzle1(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 1076)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day01.SolvePuzzle2(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 6379)
}
