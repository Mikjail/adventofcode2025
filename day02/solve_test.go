package day02_test

import (
	"spissable/advent-of-go-template/day02"
	"spissable/advent-of-go-template/utils"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day02.SolvePuzzle1(input)
	utils.LogResult(t, result)
	assert.Equal(t, 28846518423, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day02.SolvePuzzle2(input)
	utils.LogResult(t, result)
	assert.Equal(t, 31578210022, result)
}
