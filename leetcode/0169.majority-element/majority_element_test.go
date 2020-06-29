package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, 2, majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
