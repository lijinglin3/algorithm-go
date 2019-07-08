package leetcode

import (
	"reflect"
	"testing"
)

var (
	case1, target1, result1 = []int{2, 7, 11, 15}, 9, []int{0, 1}
	case2, target2, result2 = []int{7, 11, 15}, 22, []int{0, 2}
	case3, target3, result3 = []int{11, 15}, 26, []int{0, 1}
)

func TestTwoSum1(t *testing.T) {
	if !reflect.DeepEqual(twoSum1(case1, target1), result1) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum1(case2, target2), result2) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum1(case3, target3), result3) {
		t.Fail()
	}
	if twoSum1(nil, 0) != nil {
		t.Fail()
	}
}

func TestTwoSum2(t *testing.T) {
	if !reflect.DeepEqual(twoSum2(case1, target1), result1) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum2(case2, target2), result2) {
		t.Fail()
	}
	if !reflect.DeepEqual(twoSum2(case3, target3), result3) {
		t.Fail()
	}
	if twoSum2(nil, 0) != nil {
		t.Fail()
	}
}
