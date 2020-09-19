package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	mapSum := Constructor()

	mapSum.Insert("apple", 3)
	assert.Equal(t, 0, mapSum.Sum("ab"))
	assert.Equal(t, 3, mapSum.Sum("ap"))

	mapSum.Insert("app", 2)
	assert.Equal(t, 5, mapSum.Sum("ap"))
}
