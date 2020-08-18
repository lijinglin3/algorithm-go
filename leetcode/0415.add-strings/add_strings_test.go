package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	assert.Equal(t, "988888888", addStrings("1234567", "987654321"))
	assert.Equal(t, "1111111110", addStrings("123456789", "987654321"))
}
