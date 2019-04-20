package leetcode

import (
	"reflect"
	"testing"
)

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
