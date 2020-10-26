package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestKthToLast(t *testing.T) {
	assert.Equal(t, 4, kthToLast(leetcode.NewListNode("[1,2,3,4,5]"), 2))
}
