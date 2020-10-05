package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	assert.Equal(t, 3, pathSum(leetcode.NewTreeNode("[10,5,-3,3,2,null,11,3,-2,null,1]"), 8))
}
