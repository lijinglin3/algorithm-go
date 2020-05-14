package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func TestInsertIntoBST(t *testing.T) {
	assert.Equal(t, common.TreeNodeDecoder("[4, 2, 7, 1, 3, 5]"), insertIntoBST(common.TreeNodeDecoder("[4, 2, 7, 1, 3]"), 5))
	assert.Equal(t, common.TreeNodeDecoder("[4, 2, 7, 1, 3, null, 10]"), insertIntoBST(common.TreeNodeDecoder("[4, 2, 7, 1, 3]"), 10))
	assert.Equal(t, common.TreeNodeDecoder("[4, 2, 7, 1, 3, null, null, 0]"), insertIntoBST(common.TreeNodeDecoder("[4, 2, 7, 1, 3]"), 0))
	assert.Equal(t, common.TreeNodeDecoder("[4, 2, 7, 1, 3]"), insertIntoBST(common.TreeNodeDecoder("[4, 2, 7, 1, 3]"), 4))
	assert.Equal(t, common.TreeNodeDecoder("[0]"), insertIntoBST(nil, 0))
}
