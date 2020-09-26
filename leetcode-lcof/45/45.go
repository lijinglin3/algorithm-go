package leetcode

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	rst := make([]string, len(nums))
	for k, v := range nums {
		rst[k] = strconv.Itoa(v)
	}
	sort.Slice(rst, func(i, j int) bool {
		return rst[i]+rst[j] < rst[j]+rst[i]
	})
	var ret string
	for _, v := range rst {
		ret += v
	}
	return ret
}
