package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePossibleNextMoves(t *testing.T) {
	assert.ElementsMatch(t, []string{"--++", "+--+", "++--"}, generatePossibleNextMoves("++++"))
	assert.ElementsMatch(t, []string{"+---"}, generatePossibleNextMoves("+-++"))
	assert.ElementsMatch(t, []string{}, generatePossibleNextMoves("+-"))
	assert.ElementsMatch(t, []string{}, generatePossibleNextMoves("+"))
}
