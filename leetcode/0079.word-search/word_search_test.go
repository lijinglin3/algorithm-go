package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	example := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	assert.Equal(t, true, exist(example, "ABCCED"))
	assert.Equal(t, false, exist(example, "ABCB"))
}
