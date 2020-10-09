package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))
	assert.Equal(t, 2123366400, nthUglyNumber(1690))
	assert.Equal(t, 51200000, nthUglyNumber(1000))
}
