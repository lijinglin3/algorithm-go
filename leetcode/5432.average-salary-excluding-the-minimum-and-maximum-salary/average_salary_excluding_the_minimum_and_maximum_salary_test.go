package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	assert.Equal(t, 2500.0, average([]int{4000, 3000, 1000, 2000}))
}
