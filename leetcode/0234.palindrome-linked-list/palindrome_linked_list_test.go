package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	if !assert.Equal(t, true, IsPalindrome(toListNode([]int{}))) {
		t.Fail()
	}

	if !assert.Equal(t, true, IsPalindrome(toListNode([]int{1}))) {
		t.Fail()
	}

	if !assert.Equal(t, true, IsPalindrome(toListNode([]int{1, 2, 3, 4, 3, 2, 1}))) {
		t.Fail()
	}

	if !assert.Equal(t, false, IsPalindrome(toListNode([]int{1, 2}))) {
		t.Fail()
	}

	if !assert.Equal(t, false, IsPalindrome(toListNode([]int{1, 2, 3}))) {
		t.Fail()
	}

	return
}

func toListNode(values []int) (node *ListNode) {
	for _, v := range values {
		node = &ListNode{Val: v, Next: node}
	}
	return
}
