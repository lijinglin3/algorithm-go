package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	logger := Constructor()

	assert.Equal(t, true, logger.ShouldPrintMessage(1, "foo"))
	assert.Equal(t, true, logger.ShouldPrintMessage(2, "bar"))
	assert.Equal(t, false, logger.ShouldPrintMessage(3, "foo"))
	assert.Equal(t, false, logger.ShouldPrintMessage(8, "bar"))
	assert.Equal(t, false, logger.ShouldPrintMessage(10, "foo"))
	assert.Equal(t, true, logger.ShouldPrintMessage(11, "foo"))
}
