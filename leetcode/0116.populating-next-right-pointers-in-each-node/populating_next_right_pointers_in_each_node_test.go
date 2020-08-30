package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	example := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	result := connect(example)

	assert.Equal(t, true, result.Next == nil)
	assert.Equal(t, true, result.Left.Next == result.Right)

	assert.Equal(t, (*Node)(nil), connect(nil))
}
