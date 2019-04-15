package leetcode

import (
	"reflect"
	"testing"
)

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if i != j && (nums[i]+nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	if !reflect.DeepEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum([]int{2, 7, 11, 15}, 4), []int{}) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum([]int{2, 7, 11, 15}, 0), []int{}) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum([]int{}, 0), []int{}) {
		t.Fail()
	}
}