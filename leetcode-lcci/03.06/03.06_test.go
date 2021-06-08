package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	s := Constructor()

	s.Enqueue([]int{0, 0})
	s.Enqueue([]int{1, 0})

	assert.Equal(t, []int{0, 0}, s.DequeueCat())
	assert.Equal(t, []int{-1, -1}, s.DequeueDog())
	assert.Equal(t, []int{1, 0}, s.DequeueAny())

	s.Enqueue([]int{0, 0})
	s.Enqueue([]int{1, 0})
	s.Enqueue([]int{2, 1})
	s.Enqueue([]int{3, 0})
	s.Enqueue([]int{4, 1})
	s.Enqueue([]int{5, 0})
	s.Enqueue([]int{6, 1})

	assert.Equal(t, []int{2, 1}, s.DequeueDog())
	assert.Equal(t, []int{0, 0}, s.DequeueCat())
	assert.Equal(t, []int{1, 0}, s.DequeueAny())
	assert.Equal(t, []int{3, 0}, s.DequeueAny())
	assert.Equal(t, []int{4, 1}, s.DequeueAny())
	assert.Equal(t, []int{5, 0}, s.DequeueAny())
	assert.Equal(t, []int{6, 1}, s.DequeueAny())
}
