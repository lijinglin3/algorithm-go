package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHint(t *testing.T) {
	assert.Equal(t, "1A3B", getHint("1807", "7810"))
	assert.Equal(t, "1A1B", getHint("1123", "0111"))
	assert.Equal(t, "0A4B", getHint("1122", "2211"))
}
