package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	assert.Equal(t, []int{3, 9, 20, 15, 7}, levelOrder(leetcode.NewTreeNode("[3,9,20,null,null,15,7]")))
}
