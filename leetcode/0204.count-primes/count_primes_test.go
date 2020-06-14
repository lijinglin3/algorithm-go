package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimes(t *testing.T) {
	assert.Equal(t, 4, countPrimes(10))
}
