package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestReverseList(t *testing.T) {
	if !assert.ElementsMatch(t, []int{}, toSlice(reverseList(toListNode([]int{})))) {
		t.Fail()
	}

	if !assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, toSlice(reverseList(toListNode([]int{5, 4, 3, 2, 1})))) {
		t.Fail()
	}
	return
}

func TestReverseListByRecursion(t *testing.T) {
	if !assert.ElementsMatch(t, []int{}, toSlice(reverseListByRecursion(toListNode([]int{})))) {
		t.Fail()
	}

	if !assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, toSlice(reverseListByRecursion(toListNode([]int{5, 4, 3, 2, 1})))) {
		t.Fail()
	}
	return
}

func toSlice(node *leetcode.ListNode) (values []int) {
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return
}

func toListNode(values []int) (node *leetcode.ListNode) {
	for _, v := range values {
		node = &leetcode.ListNode{Val: v, Next: node}
	}
	return
}
