package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "1001010", addBinary("101110", "11100"))
}
