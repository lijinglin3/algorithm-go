package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {
	assert.Contains(t, []string{"eert", "eetr"}, frequencySort("tree"))
	assert.Contains(t, []string{"bbaA", "bbAa"}, frequencySort("Aabb"))
}
