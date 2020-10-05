package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin3/algorithm-go/leetcode"
)

func TestIsPalindrome(t *testing.T) {
	if !assert.Equal(t, true, isPalindrome(toListNode([]int{}))) {
		t.Fail()
	}

	if !assert.Equal(t, true, isPalindrome(toListNode([]int{1}))) {
		t.Fail()
	}

	if !assert.Equal(t, true, isPalindrome(toListNode([]int{1, 2, 3, 4, 3, 2, 1}))) {
		t.Fail()
	}

	if !assert.Equal(t, false, isPalindrome(toListNode([]int{1, 2}))) {
		t.Fail()
	}

	if !assert.Equal(t, false, isPalindrome(toListNode([]int{1, 2, 3}))) {
		t.Fail()
	}
}

func toListNode(values []int) (node *leetcode.ListNode) {
	for _, v := range values {
		node = &leetcode.ListNode{Val: v, Next: node}
	}
	return
}
