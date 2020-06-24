package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUgly(t *testing.T) {
	assert.Equal(t, true, isUgly(30))
	assert.Equal(t, false, isUgly(14))
	assert.Equal(t, false, isUgly(0))
	assert.Equal(t, false, isUgly(-1))
}
