package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	cases := []*leetcode.TreeNode{
		leetcode.NewTreeNode("[1, 2, 3, 4, 5]"),
		leetcode.NewTreeNode("[1, 2, 2, 3, null, 4]"),
		leetcode.NewTreeNode("[1, 2, 2, 3, 4, 4, 3]"),
		{},
		nil,
	}
	results := []bool{
		false,
		false,
		true,
		true,
		true,
	}

	for i := range cases {
		assert.Equal(t, results[i], isSymmetric(cases[i]))
	}
}
