package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListToTreeNode(t *testing.T) {
	cases := []string{
		"[]",
		"[1, 2, 3, 4, 5]",
		"[1, 2, 2, 3, null, 4]",
		"[1, 2, 2, 3, 4, 4, 3]",
		"[1, null, 2, null, 3]",
	}
	results := []*TreeNode{
		nil,

		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},

		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},

		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},

		{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
	}

	for i := range cases {
		assert.Equal(t, results[i], TreeNodeDecoder(cases[i]))
	}
}
