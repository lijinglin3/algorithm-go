package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	example := []int{1, 2, 3, 1}

	assert.Equal(t, true, containsNearbyAlmostDuplicate(example, 3, 0))
	assert.Equal(t, false, containsNearbyAlmostDuplicate(example, 2, 0))
}
