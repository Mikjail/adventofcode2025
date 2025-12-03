package day03_test

import (
	"spissable/advent-of-go-template/day03"
	"spissable/advent-of-go-template/utils"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day03.SolvePuzzle1(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 17034)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day03.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
