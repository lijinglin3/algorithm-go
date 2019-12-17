package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	if !assert.ElementsMatch(t, []int{}, toSlice(ReverseList(toListNode([]int{})))) {
		t.Fail()
	}

	if !assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, toSlice(ReverseList(toListNode([]int{5, 4, 3, 2, 1})))) {
		t.Fail()
	}
	return
}

func toSlice(node *ListNode) (values []int) {
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return
}

func toListNode(values []int) (node *ListNode) {
	for _, v := range values {
		node = &ListNode{Val: v, Next: node}
	}
	return
}
