package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	queue := Constructor(3)

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Rear())
	assert.Equal(t, true, queue.EnQueue(1))
	assert.Equal(t, true, queue.EnQueue(2))
	assert.Equal(t, true, queue.EnQueue(3))
	assert.Equal(t, false, queue.EnQueue(4))
	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Rear())
	assert.Equal(t, true, queue.DeQueue())
	assert.Equal(t, true, queue.EnQueue(4))
	assert.Equal(t, 4, queue.Rear())
	assert.Equal(t, true, queue.DeQueue())
	assert.Equal(t, true, queue.DeQueue())
	assert.Equal(t, true, queue.DeQueue())
	assert.Equal(t, false, queue.DeQueue())
}
