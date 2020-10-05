package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestIsSubStructure(t *testing.T) {
	assert.Equal(t, true, isSubStructure(leetcode.NewTreeNode("[3,4,5,1,2]"), leetcode.NewTreeNode("[4,1]")))
	assert.Equal(t, false, isSubStructure(leetcode.NewTreeNode("[1,2,3]"), leetcode.NewTreeNode("[3,1]")))
}
