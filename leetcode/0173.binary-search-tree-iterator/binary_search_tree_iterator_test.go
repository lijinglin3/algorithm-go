package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestConstructor(t *testing.T) {
	assert.Equal(t, BSTIterator{}, Constructor((*leetcode.TreeNode)(nil)))
	assert.Equal(t, BSTIterator{stack: []*leetcode.TreeNode{{}}, length: 1}, Constructor(&leetcode.TreeNode{}))

	constructor := Constructor(leetcode.NewTreeNode("[7, 3, 15, null, null, 9, 20]"))
	for _, v := range []int{3, 7, 9, 15, 20} {
		assert.Equal(t, true, constructor.HasNext())
		assert.Equal(t, v, constructor.Next())
	}
	assert.Equal(t, false, constructor.HasNext())
}
