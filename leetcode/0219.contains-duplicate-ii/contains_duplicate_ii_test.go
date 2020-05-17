package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	example := []int{1, 2, 3, 1}
	cases := []int{3, 2}
	results := []bool{true, false}

	for i, v := range cases {
		assert.Equal(t, results[i], containsNearbyDuplicate(example, v))
		assert.Equal(t, results[i], containsNearbyDuplicate2(example, v))
	}
}
