package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestReverseKGroup(t *testing.T) {
	list := "[1, 2, 3, 4, 5, 6, 7, 8]"

	assert.Equal(t, leetcode.ListNodeDecoder(list),
		reverseKGroup(leetcode.ListNodeDecoder(list), 1))
	assert.Equal(t, leetcode.ListNodeDecoder("[2, 1, 4, 3, 6, 5, 8, 7]"),
		reverseKGroup(leetcode.ListNodeDecoder(list), 2))
	assert.Equal(t, leetcode.ListNodeDecoder("[3, 2, 1, 6, 5, 4, 7, 8]"),
		reverseKGroup(leetcode.ListNodeDecoder(list), 3))
}
