package binary_search

import (
	"testing"
)

var (
	testSlice0 = make([]int, 0)
	testSlice1 = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	testSlice2 = []int{10, 30, 30, 30, 50, 80, 80, 80, 90, 100}
)

func TestBinarySearch(t *testing.T) {
	if BinarySearch(testSlice0, 0) != -1 {
		t.Fail()
	}
	if BinarySearch(testSlice1, 0) != -1 {
		t.Fail()
	}
	if BinarySearch(testSlice1, 30) != 2 {
		t.Fail()
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	if BinarySearchRecursive(testSlice0, 0) != -1 {
		t.Fail()
	}
	if BinarySearchRecursive(testSlice1, 0) != -1 {
		t.Fail()
	}
	if BinarySearchRecursive(testSlice1, 30) != 2 {
		t.Fail()
	}
}

func TestBinarySearchFirst(t *testing.T) {
	if BinarySearchFirst(testSlice0, 0) != -1 {
		t.Fail()
	}
	if BinarySearchFirst(testSlice2, 0) != -1 {
		t.Fail()
	}
	if BinarySearchFirst(testSlice2, 30) != 1 {
		t.Fail()
	}
	if BinarySearchFirst(testSlice2, 80) != 5 {
		t.Fail()
	}
}

func TestBinarySearchLast(t *testing.T) {
	if BinarySearchLast(testSlice0, 0) != -1 {
		t.Fail()
	}
	if BinarySearchLast(testSlice2, 0) != -1 {
		t.Fail()
	}
	if BinarySearchLast(testSlice2, 30) != 3 {
		t.Fail()
	}
	if BinarySearchLast(testSlice2, 80) != 7 {
		t.Fail()
	}
}
