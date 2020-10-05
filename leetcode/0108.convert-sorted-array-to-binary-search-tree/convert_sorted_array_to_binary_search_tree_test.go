package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestSortedArrayToBST(t *testing.T) {
	assert.Equal(t, leetcode.NewTreeNode("[0, -3, 9, -10, null, 5]"),
		sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
