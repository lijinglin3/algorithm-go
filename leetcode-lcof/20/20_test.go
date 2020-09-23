package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, true, isNumber("+123"))
	assert.Equal(t, true, isNumber("5e2"))
	assert.Equal(t, true, isNumber("3.1415926"))
	assert.Equal(t, true, isNumber("-1E-16"))
	assert.Equal(t, false, isNumber("12e"))
	assert.Equal(t, false, isNumber("1a3.14"))
	assert.Equal(t, false, isNumber("1.2.3"))
	assert.Equal(t, false, isNumber("+-5"))
	assert.Equal(t, false, isNumber("12e3.14"))
	assert.Equal(t, false, isNumber("12 3.14"))
}
