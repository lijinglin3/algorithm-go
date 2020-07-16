package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPerfectSquare(t *testing.T) {
	assert.Equal(t, true, isPerfectSquare(16))
	assert.Equal(t, false, isPerfectSquare(15))
	assert.Equal(t, false, isPerfectSquare(20))
}
