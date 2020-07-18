package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	assert.Equal(t, true, exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	assert.Equal(t, false, exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "abcd"))
}
