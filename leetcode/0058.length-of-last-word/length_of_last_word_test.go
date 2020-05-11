package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	cases := []string{
		"Hello World",
		"Hello World  ",
		"HelloWorld",
		"",
	}
	results := []int{5, 5, 10, 0}

	for i := range cases {
		assert.Equal(t, results[i], LengthOfLastWord(cases[i]))
	}
}
