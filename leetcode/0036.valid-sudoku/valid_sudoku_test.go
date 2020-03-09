package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	cases := [][][]byte{
		{
			{byte(5), byte(3), byte('.'), byte('.'), byte(7), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte(6), byte('.'), byte('.'), byte(1), byte(9), byte(5), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte(9), byte(8), byte('.'), byte('.'), byte('.'), byte('.'), byte(6), byte('.')},
			{byte(8), byte('.'), byte('.'), byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte(3)},
			{byte(4), byte('.'), byte('.'), byte(8), byte('.'), byte(3), byte('.'), byte('.'), byte(1)},
			{byte(7), byte('.'), byte('.'), byte('.'), byte(2), byte('.'), byte('.'), byte('.'), byte(6)},
			{byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte('.'), byte(2), byte(8), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte(4), byte(1), byte(9), byte('.'), byte('.'), byte(5)},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte(8), byte('.'), byte('.'), byte(7), byte(9)},
		},
		{
			{byte(8), byte(3), byte('.'), byte('.'), byte(7), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte(6), byte('.'), byte('.'), byte(1), byte(9), byte(5), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte(9), byte(8), byte('.'), byte('.'), byte('.'), byte('.'), byte(6), byte('.')},
			{byte(8), byte('.'), byte('.'), byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte(3)},
			{byte(4), byte('.'), byte('.'), byte(8), byte('.'), byte(3), byte('.'), byte('.'), byte(1)},
			{byte(7), byte('.'), byte('.'), byte('.'), byte(2), byte('.'), byte('.'), byte('.'), byte(6)},
			{byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte('.'), byte(2), byte(8), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte(4), byte(1), byte(9), byte('.'), byte('.'), byte(5)},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte(8), byte('.'), byte('.'), byte(7), byte(9)},
		},
		{
			{byte(5), byte(3), byte('.'), byte('.'), byte(7), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte(6), byte('.'), byte('.'), byte(1), byte(9), byte(5), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte(9), byte(8), byte('.'), byte('.'), byte('.'), byte('.'), byte(6), byte('.')},
			{byte(8), byte('.'), byte('.'), byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte(3)},
			{byte(4), byte('.'), byte('.'), byte(8), byte('.'), byte(3), byte('.'), byte('.'), byte(1)},
			{byte(7), byte('.'), byte('.'), byte('.'), byte(2), byte('.'), byte('.'), byte('.'), byte(6)},
			{byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte('.'), byte(2), byte(8), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte(4), byte(1), byte(9), byte('.'), byte('.'), byte(5)},
			{byte(7), byte('.'), byte('.'), byte('.'), byte(8), byte('.'), byte('.'), byte(7), byte(9)},
		},
		{
			{byte(5), byte(3), byte('.'), byte('.'), byte(7), byte('.'), byte('.'), byte('.'), byte(3)},
			{byte(6), byte('.'), byte('.'), byte(1), byte(9), byte(5), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte(9), byte(8), byte('.'), byte('.'), byte('.'), byte('.'), byte(6), byte('.')},
			{byte(8), byte('.'), byte('.'), byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte(3)},
			{byte(4), byte('.'), byte('.'), byte(8), byte('.'), byte(3), byte('.'), byte('.'), byte(1)},
			{byte(7), byte('.'), byte('.'), byte('.'), byte(2), byte('.'), byte('.'), byte('.'), byte(6)},
			{byte('.'), byte(6), byte('.'), byte('.'), byte('.'), byte('.'), byte(2), byte(8), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte(4), byte(1), byte(9), byte('.'), byte('.'), byte(5)},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte(8), byte('.'), byte('.'), byte(7), byte(9)},
		},
	}
	results := []bool{true, false, false, false}

	for i := range cases {
		assert.Equal(t, IsValidSudoku(cases[i]), results[i])
	}
}
