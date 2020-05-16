package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestMergeTwoLists(t *testing.T) {
	result := "[1, 2, 3, 4, 5, 6, 7]"
	listNode1 := "[1, 3, 5, 7]"
	listNode2 := "[2, 4, 6]"
	assert.Equal(t, leetcode.ListNodeDecoder(result),
		mergeTwoLists(
			leetcode.ListNodeDecoder(listNode1),
			leetcode.ListNodeDecoder(listNode2),
		),
	)
	assert.Equal(t, leetcode.ListNodeDecoder(result),
		mergeTwoLists(
			leetcode.ListNodeDecoder(listNode2),
			leetcode.ListNodeDecoder(listNode1),
		),
	)
}
