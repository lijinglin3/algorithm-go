package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestOddEvenList(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[]"), oddEvenList(leetcode.NewListNode("[]")))
	assert.Equal(t, leetcode.NewListNode("[1,3,5,2,4]"), oddEvenList(leetcode.NewListNode("[1,2,3,4,5]")))
}
