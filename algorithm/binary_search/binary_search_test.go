package binarysearch

import (
	"testing"
)

var (
	testSlice0 = make([]int, 0)
	testSlice1 = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	testSlice2 = []int{10, 30, 30, 30, 50, 80, 80, 80, 90, 100}
)

func TestBinarySearch(t *testing.T) {
	if Search(testSlice0, 0) != -1 {
		t.Fail()
	}
	if Search(testSlice1, 0) != -1 {
		t.Fail()
	}
	if Search(testSlice1, 30) != 2 {
		t.Fail()
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	if SearchRecursive(testSlice0, 0) != -1 {
		t.Fail()
	}
	if SearchRecursive(testSlice1, 0) != -1 {
		t.Fail()
	}
	if SearchRecursive(testSlice1, 30) != 2 {
		t.Fail()
	}
}

func TestBinarySearchFirst(t *testing.T) {
	if SearchFirst(testSlice0, 0) != -1 {
		t.Fail()
	}
	if SearchFirst(testSlice2, 0) != -1 {
		t.Fail()
	}
	if SearchFirst(testSlice2, 30) != 1 {
		t.Fail()
	}
	if SearchFirst(testSlice2, 80) != 5 {
		t.Fail()
	}
}

func TestBinarySearchLast(t *testing.T) {
	if SearchLast(testSlice0, 0) != -1 {
		t.Fail()
	}
	if SearchLast(testSlice2, 0) != -1 {
		t.Fail()
	}
	if SearchLast(testSlice2, 30) != 3 {
		t.Fail()
	}
	if SearchLast(testSlice2, 80) != 7 {
		t.Fail()
	}
}
