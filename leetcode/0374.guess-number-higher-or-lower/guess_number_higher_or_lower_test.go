package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessNumber(t *testing.T) {
	pick = 6
	assert.Equal(t, pick, guessNumber(10))
	pick = 20
	assert.Equal(t, -1, guessNumber(10))
}
