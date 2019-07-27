package binary_search

// BinarySearch 二分查找
func BinarySearch(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] == v:
			return mid
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchRecursive 使用递归方式的二分查找
func BinarySearchRecursive(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	return bs(nums, v, 0, length-1)
}

func bs(nums []int, v int, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	switch {
	case nums[mid] == v:
		return mid
	case nums[mid] > v:
		return bs(nums, v, low, mid-1)
	case nums[mid] < v:
		return bs(nums, v, mid+1, high)
	default:
		return -1
	}
}

// BinarySearchFirst 查找第一个等于给定值的元素
func BinarySearchFirst(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) >> 1
		switch {
		case nums[mid] == v:
			if mid == 0 || nums[mid-1] != v {
				return mid
			}
			high = mid - 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}
	return -1
}

// BinarySearchLast 查找最后一个值等于给定值的元素
func BinarySearchLast(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) >> 1

		switch {
		case nums[mid] == v:
			if mid == length-1 || nums[mid+1] != v {
				return mid
			}
			low = mid + 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchFirstGT 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1

	for low <= high {
		mid := (low + high) >> 1

		switch {
		case nums[mid] == v:
			if mid != length-1 && nums[mid+1] > v {
				return mid + 1
			}
			low = mid + 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	for low <= high {
		mid := (high + low) >> 1
		if nums[mid] > v {
			high = mid - 1
		} else if nums[mid] < v {
			low = mid + 1
		} else {

		}
	}

	return -1
}

// BinarySearchLastLT 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(nums []int, v int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	low, high := 0, length-1
	for low <= high {
		mid := (low + high) >> 1

		switch {
		case nums[mid] == v:
			if mid != 0 && nums[mid-1] < v {
				return mid - 1
			}
			high = mid - 1
		case nums[mid] > v:
			high = mid - 1
		case nums[mid] < v:
			low = mid + 1
		}
	}

	return -1
}
