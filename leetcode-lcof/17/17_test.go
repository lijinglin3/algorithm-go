package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintNumbers(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, printNumbers(1))
}
