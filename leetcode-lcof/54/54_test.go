package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	assert.Equal(t, 4, kthLargest(leetcode.NewTreeNode("[3,1,4,null,2]"), 1))
	assert.Equal(t, 3, kthLargest(leetcode.NewTreeNode("[3,1,4,null,2]"), 2))
	assert.Equal(t, 2, kthLargest(leetcode.NewTreeNode("[3,1,4,null,2]"), 3))
	assert.Equal(t, 1, kthLargest(leetcode.NewTreeNode("[3,1,4,null,2]"), 4))
	assert.Equal(t, 1, kthLargest(leetcode.NewTreeNode("[3,1,4,null,2]"), 5))
}
