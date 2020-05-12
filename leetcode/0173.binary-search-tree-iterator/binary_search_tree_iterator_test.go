package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestConstructor(t *testing.T) {
	assert.Equal(t, BSTIterator{}, Constructor((*TreeNode)(nil)))
	assert.Equal(t, BSTIterator{stack: []*TreeNode{{}}, length: 1}, Constructor(&TreeNode{}))

	constructor := Constructor(TreeNodeDecoder("[7, 3, 15, null, null, 9, 20]"))
	for _, v := range []int{3, 7, 9, 15, 20} {
		assert.Equal(t, true, constructor.HasNext())
		assert.Equal(t, v, constructor.Next())
	}
	assert.Equal(t, false, constructor.HasNext())
}
