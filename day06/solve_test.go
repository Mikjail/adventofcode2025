package day06_test

import (
	"spissable/advent-of-go-template/day06"
	"spissable/advent-of-go-template/utils"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day06.SolvePuzzle1(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 4648618073226)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day06.SolvePuzzle2(input)
	utils.LogResult(t, result)
	assert.Equal(t, result, 7329921182115)
}
