package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	assert.Equal(t, true, isSubtree(leetcode.NewTreeNode("[3,4,5,1,2]"), leetcode.NewTreeNode("[4,1,2]")))
	assert.Equal(t, false, isSubtree(leetcode.NewTreeNode("[3,4,5,1,2,null,null,null,null,0]"), leetcode.NewTreeNode("[4,1,2]")))
}
