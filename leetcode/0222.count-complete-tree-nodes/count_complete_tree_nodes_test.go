package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestCountNodes(t *testing.T) {
	assert.Equal(t, 0, countNodes(leetcode.NewTreeNode("[]")))
	assert.Equal(t, 1, countNodes(leetcode.NewTreeNode("[1]")))
	assert.Equal(t, 6, countNodes(leetcode.NewTreeNode("[1,2,3,4,5,6]")))
}
