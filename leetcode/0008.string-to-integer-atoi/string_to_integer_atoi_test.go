package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, -43, myAtoi("    -43   22"))
	assert.Equal(t, -12, myAtoi("-12+34"))
	assert.Equal(t, 0, myAtoi("abc"))
}
