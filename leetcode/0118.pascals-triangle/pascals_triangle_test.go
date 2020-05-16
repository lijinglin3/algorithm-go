package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	result := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	assert.Equal(t, result, generate(5))
	assert.Equal(t, [][]int{}, generate(0))
}
