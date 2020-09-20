package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	assert.Equal(t,
		leetcode.NewTreeNode("[3,9,20,null,null,15,7]"),
		buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}),
	)
}
