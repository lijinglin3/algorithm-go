package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T) {
	assert.Equal(t, [][]int(nil), zigzagLevelOrder(nil))
	assert.Equal(t,
		[][]int{{3}, {20, 9}, {15, 7}},
		zigzagLevelOrder(leetcode.NewTreeNode("[3,9,20,null,null,15,7]")),
	)
}
