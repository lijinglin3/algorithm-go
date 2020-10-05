package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestGetKthFromEnd(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[4,5]"), getKthFromEnd(leetcode.NewListNode("[1,2,3,4,5]"), 2))
	assert.Equal(t, leetcode.NewListNode("[1,2,3,4,5]"), getKthFromEnd(leetcode.NewListNode("[1,2,3,4,5]"), 7))
}
