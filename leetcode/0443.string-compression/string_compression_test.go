package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	assert.Equal(t, 1, compress([]byte("a")))
	assert.Equal(t, 4, compress([]byte("abbbbbbbbbbbb")))
}
