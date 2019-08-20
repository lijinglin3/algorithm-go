package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testSlices := [][]int{
		{},
		{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
		{1, 3, 5, 7, 9, 1, 3, 5, 7, 9},
	}
	resultSlices := [][]int{
		{},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 1, 3, 3, 5, 5, 7, 7, 9, 9},
	}

	for i, testSlice := range testSlices {
		testNums := make([]int, len(testSlice))

		copy(testNums, testSlice)
		BubbleSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		BucketSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		BucketSortSimple(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		CountingSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		InsertionSort(testNums, len(testNums))
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		MergeSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		QuickSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}

		copy(testNums, testSlice)
		SelectionSort(testNums)
		if !reflect.DeepEqual(testNums, resultSlices[i]) {
			t.Fail()
		}
	}
}
