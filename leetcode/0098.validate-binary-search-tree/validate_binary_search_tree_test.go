package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode"
)

func TestIsValidBST(t *testing.T) {
	assert.Equal(t, true, isValidBST(leetcode.NewTreeNode("[2, 1, 3]")))
	assert.Equal(t, false,
		isValidBST(leetcode.NewTreeNode("[5, 1, 4, null, null, 3, 6]")))
}
