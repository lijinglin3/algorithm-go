package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNodeDecoder(t *testing.T) {
	assert.Panics(t, func() { NewListNode("") })

	cases := []string{
		"[]",
		"[1, 2, 3, 4, 5]",
	}
	results := []*ListNode{
		nil,
		{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	for i := range cases {
		assert.Equal(t, results[i], NewListNode(cases[i]))
	}
}

func TestListNode_Last(t *testing.T) {
	list := NewListNode("[1, 2, 3, 4, 5]")
	assert.Equal(t, NewListNode("[5]"), list.Last())
	assert.Equal(t, (*ListNode)(nil), (*ListNode)(nil).Last())
}

func TestTreeNodeDecoder(t *testing.T) {
	assert.Panics(t, func() { NewTreeNode("") })

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
		assert.Equal(t, results[i], NewTreeNode(cases[i]))
	}
}
