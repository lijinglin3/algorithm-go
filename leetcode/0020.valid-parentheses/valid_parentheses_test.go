package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	cases := []string{
		"[]{}()",
		"([{]})",
		"",
	}
	results := []bool{
		true,
		false,
		true,
	}

	for i, v := range cases {
		assert.Equal(t, results[i], isValid(v))
	}
}
