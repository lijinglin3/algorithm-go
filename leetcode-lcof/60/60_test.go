package leetcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []float64{
		0.16667,
		0.16667,
		0.16667,
		0.16667,
		0.16667,
		0.16667,
	}, format(twoSum(1)))
	assert.Equal(t, []float64{
		0.02778,
		0.05556,
		0.08333,
		0.11111,
		0.13889,
		0.16667,
		0.13889,
		0.11111,
		0.08333,
		0.05556,
		0.02778,
	}, format(twoSum(2)))
	assert.Equal(t, []float64{
		0.00463,
		0.01389,
		0.02778,
		0.04630,
		0.06944,
		0.09722,
		0.11574,
		0.12500,
		0.12500,
		0.11574,
		0.09722,
		0.06944,
		0.04630,
		0.02778,
		0.01389,
		0.00463,
	}, format(twoSum(3)))
}

func format(nums []float64) []float64 {
	for i := range nums {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", nums[i]), 10)
		nums[i] = v
	}
	return nums
}
