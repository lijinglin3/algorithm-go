package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Equal(t, leetcode.NewListNode("[1]"), deleteDuplicates(leetcode.NewListNode("[1]")))
	assert.Equal(t, leetcode.NewListNode("[1,2,5]"), deleteDuplicates(leetcode.NewListNode("[1,2,3,3,4,4,5]")))
}
