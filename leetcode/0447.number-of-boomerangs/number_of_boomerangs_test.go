package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfBoomerangs(t *testing.T) {
	assert.Equal(t, 2, numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
}
