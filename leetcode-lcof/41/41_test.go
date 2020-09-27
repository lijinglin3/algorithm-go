package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()

	c.AddNum(1)
	assert.Equal(t, 1.0, c.FindMedian())
	c.AddNum(2)
	assert.Equal(t, 1.5, c.FindMedian())

	c.AddNum(3)
	assert.Equal(t, 2.0, c.FindMedian())

	c.AddNum(4)
	c.AddNum(5)
	assert.Equal(t, 3.0, c.FindMedian())
}
