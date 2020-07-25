package leetcode

const max = 2147483647

// ArrayReader ...
type ArrayReader struct {
	data []int
}

func (r *ArrayReader) get(index int) int {
	if index >= len(r.data) {
		return max
	}
	return r.data[index]
}

func search(reader ArrayReader, target int) int {
	left, right := 0, 1
	for reader.get(right) != max {
		right <<= 1
	}
	for left <= right {
		mid := (left + right) / 2
		m := reader.get(mid)
		switch {
		case target == m:
			return mid
		case target > m:
			left = mid + 1
		case target < m:
			right = mid - 1
		}
	}
	return -1
}
