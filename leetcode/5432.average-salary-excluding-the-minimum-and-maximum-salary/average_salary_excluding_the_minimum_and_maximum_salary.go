package leetcode

import "math"

func average(salary []int) float64 {
	length, sum, min, max := len(salary), 0, math.MaxInt64, 0
	for i := 0; i < length; i++ {
		sum += salary[i]
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
	}
	return float64(sum-min-max) / float64(length-2)
}
