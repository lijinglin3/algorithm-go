package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPostorder(t *testing.T) {
	assert.Equal(t, false, verifyPostorder([]int{1, 6, 3, 2, 5}))
	assert.Equal(t, true, verifyPostorder([]int{1, 3, 2, 6, 5}))
}
