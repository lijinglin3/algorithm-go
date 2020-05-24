package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	assert.Equal(t, (*leetcode.ListNode)(nil), rotateRight(leetcode.NewListNode("[]"), 10))
	assert.Equal(t, leetcode.NewListNode("[3,4,5,1,2]"), rotateRight(leetcode.NewListNode("[1,2,3,4,5]"), 3))
}
