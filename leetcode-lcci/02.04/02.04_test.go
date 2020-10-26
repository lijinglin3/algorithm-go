package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[1,2,2,4,3,5]"), partition(leetcode.NewListNode("[1,4,3,2,5,2]"), 3))
}
