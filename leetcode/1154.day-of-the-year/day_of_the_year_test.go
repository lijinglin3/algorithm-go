package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOfYear(t *testing.T) {
	assert.Equal(t, 9, dayOfYear("2019-01-09"))
	assert.Equal(t, 41, dayOfYear("2019-02-10"))
	assert.Equal(t, 60, dayOfYear("2003-03-01"))
	assert.Equal(t, 61, dayOfYear("2004-03-01"))
}
