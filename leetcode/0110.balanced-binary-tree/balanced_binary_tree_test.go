package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

/*
1
 2
  3
   4
*/
func TestIsBalanced(t *testing.T) {
	cases := []*leetcode.TreeNode{
		{},
		nil,
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		leetcode.NewTreeNode("[1, 2, 2, 3, null, 4]"),
		leetcode.NewTreeNode("[1, null, 2, null, 3]"),
		leetcode.NewTreeNode("[1, null, 2, null, 3, null, 4]"),
		leetcode.NewTreeNode("[1, 2, null, 3, null, 4]"),
	}
	results := []bool{
		true,
		true,
		true,
		true,
		false,
		false,
		false,
	}

	for i := range cases {
		assert.Equal(t, results[i], isBalanced(cases[i]))
	}
}
