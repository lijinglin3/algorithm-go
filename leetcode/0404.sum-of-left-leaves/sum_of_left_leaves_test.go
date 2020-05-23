package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	assert.Equal(t, 0, sumOfLeftLeaves(leetcode.NewTreeNode("[]")))
	assert.Equal(t, 24, sumOfLeftLeaves(leetcode.NewTreeNode("[3,9,20,null,null,15,7]")))
}
