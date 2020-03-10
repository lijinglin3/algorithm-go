package leetcode

import (
	"testing"

	. "github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	cases := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		{},
		nil,
	}
	results := []int{3, 0, 0}

	for i := range cases {
		assert.Equal(t, results[i], DiameterOfBinaryTree(cases[i]))
	}
}
