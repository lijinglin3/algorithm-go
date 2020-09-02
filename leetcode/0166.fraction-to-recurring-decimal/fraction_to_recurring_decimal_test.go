package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractionToDecimal(t *testing.T) {
	assert.Equal(t, "NAN", fractionToDecimal(1, 0))
	assert.Equal(t, "2", fractionToDecimal(4, 2))
	assert.Equal(t, "0.5", fractionToDecimal(1, 2))
	assert.Equal(t, "0.(6)", fractionToDecimal(2, 3))
	assert.Equal(t, "-0.(6)", fractionToDecimal(-2, 3))
}
