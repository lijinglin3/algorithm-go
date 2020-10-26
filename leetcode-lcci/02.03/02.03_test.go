package leetcode

import (
	"testing"

	"github.com/lijinglin3/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	example := leetcode.NewListNode("[4,5,1,9]")
	deleteNode(example.Next)
	assert.Equal(t, leetcode.NewListNode("[4,1,9]"), example)
}
