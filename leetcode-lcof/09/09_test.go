package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	cq := Constructor()

	assert.Equal(t, -1, cq.DeleteHead())
	assert.Equal(t, -1, cq.DeleteHead())

	cq.AppendTail(1)
	cq.AppendTail(2)
	cq.AppendTail(3)

	assert.Equal(t, 1, cq.DeleteHead())
	assert.Equal(t, 2, cq.DeleteHead())
	assert.Equal(t, 3, cq.DeleteHead())
	assert.Equal(t, -1, cq.DeleteHead())
	assert.Equal(t, -1, cq.DeleteHead())
}
