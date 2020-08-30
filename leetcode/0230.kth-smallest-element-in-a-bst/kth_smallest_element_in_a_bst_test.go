package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	assert.Equal(t, 4, kthSmallest(leetcode.NewTreeNode("[3,1,4,null,2]"), 4))
	assert.Equal(t, -1, kthSmallest(leetcode.NewTreeNode("[3,1,4,null,2]"), 5))
}
