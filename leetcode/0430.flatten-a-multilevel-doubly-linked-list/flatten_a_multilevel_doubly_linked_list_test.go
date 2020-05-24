package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	list := &Node{}
	child := newList(1, 3)
	list.Child = child
	for child.Next != nil {
		child = child.Next
	}
	child.Child = newList(4, 6)
	nl := newList(7, 9)
	list.Next, nl.Prev = nl, list

	assert.Equal(t, newList(0, 9), flatten(list))
}

func newList(min, max int) *Node {
	list := &Node{Val: min}
	n := list
	for i := min + 1; i <= max; i++ {
		n.Next = &Node{
			Val:  i,
			Prev: n,
		}
		n = n.Next
	}
	return list
}
