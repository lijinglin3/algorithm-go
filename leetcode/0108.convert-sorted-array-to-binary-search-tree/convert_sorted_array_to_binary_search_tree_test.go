package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestSortedArrayToBST(t *testing.T) {
	assert.Equal(t, leetcode.TreeNodeDecoder("[0, -3, 9, -10, null, 5]"),
		sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
