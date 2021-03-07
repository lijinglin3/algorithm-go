package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	testdata = []bool{false, false, false, true, true}
	assert.Equal(t, 4, firstBadVersion(5))
}
