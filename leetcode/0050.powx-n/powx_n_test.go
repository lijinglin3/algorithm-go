package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	assert.Equal(t, 1/1024.0, myPow(2, -10))
	assert.Equal(t, 1.0, myPow(2, 0))
	assert.Equal(t, 1024.0, myPow(2, 10))
	assert.Equal(t, 9.261000000000001, myPow(2.1, 3))
}
