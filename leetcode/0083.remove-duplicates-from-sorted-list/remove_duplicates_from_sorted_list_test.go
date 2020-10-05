package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[]"), deleteDuplicates(leetcode.NewListNode("[]")))
	assert.Equal(t, leetcode.NewListNode("[1]"), deleteDuplicates(leetcode.NewListNode("[1]")))
	assert.Equal(t, leetcode.NewListNode("[1]"), deleteDuplicates(leetcode.NewListNode("[1,1,1]")))
	assert.Equal(t, leetcode.NewListNode("[1,2,3,4,5]"), deleteDuplicates(leetcode.NewListNode("[1,1,2,2,3,4,5,5]")))
}
